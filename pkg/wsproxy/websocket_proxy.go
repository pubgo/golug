package wsproxy

import (
	"bufio"
	"github.com/pubgo/funk/log"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024 * 10
)

// MethodOverrideParam defines the special URL parameter that is translated into the subsequent proxied streaming http request's method.
//
// Deprecated: it is preferable to use the Options parameters to WebSocketProxy to supply parameters.
var MethodOverrideParam = "method"

// TokenCookieName defines the cookie name that is translated to an 'Authorization: Bearer' header in the streaming http request's headers.
//
// Deprecated: it is preferable to use the Options parameters to WebSocketProxy to supply parameters.
var TokenCookieName = "token"

// RequestMutatorFunc can supply an alternate outgoing request.
type RequestMutatorFunc func(incoming *http.Request, outgoing *http.Request) *http.Request

// Proxy provides websocket transport upgrade to compatible endpoints.
type Proxy struct {
	h                   http.Handler
	logger              Logger
	methodOverrideParam string
	tokenCookieName     string
	requestMutator      RequestMutatorFunc
}

// Logger collects log messages.
type Logger interface {
	Warnln(...interface{})
	Debugln(...interface{})
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !websocket.IsWebSocketUpgrade(r) {
		p.h.ServeHTTP(w, r)
		return
	}
	p.proxy(w, r)
}

// Option allows customization of the proxy.
type Option func(*Proxy)

// WithMethodParamOverride allows specification of the special http parameter that is used in the proxied streaming request.
func WithMethodParamOverride(param string) Option {
	return func(p *Proxy) {
		p.methodOverrideParam = param
	}
}

// WithTokenCookieName allows specification of the cookie that is supplied as an upstream 'Authorization: Bearer' http header.
func WithTokenCookieName(param string) Option {
	return func(p *Proxy) {
		p.tokenCookieName = param
	}
}

// WithRequestMutator allows a custom RequestMutatorFunc to be supplied.
func WithRequestMutator(fn RequestMutatorFunc) Option {
	return func(p *Proxy) {
		p.requestMutator = fn
	}
}

// WithLogger allows a custom FieldLogger to be supplied
func WithLogger(logger Logger) Option {
	return func(p *Proxy) {
		p.logger = logger
	}
}

// WebsocketProxy attempts to expose the underlying handler as a bidi websocket stream with newline-delimited
// JSON as the content encoding.
//
// The HTTP Authorization header is either populated from the Sec-Websocket-Protocol field or by a cookie.
// The cookie name is specified by the TokenCookieName value.
//
// example:
//
//	Sec-Websocket-Protocol: Bearer, foobar
//
// is converted to:
//
//	Authorization: Bearer foobar
//
// Method can be overwritten with the MethodOverrideParam get parameter in the requested URL
func WebsocketProxy(h http.Handler, opts ...Option) http.Handler {
	p := &Proxy{
		h:                   h,
		logger:              logrus.New(),
		methodOverrideParam: MethodOverrideParam,
		tokenCookieName:     TokenCookieName,
	}
	for _, o := range opts {
		o(p)
	}
	return p
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func isClosedConnError(err error) bool {
	str := err.Error()
	if strings.Contains(str, "use of closed network connection") {
		return true
	}
	return websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway)
}

func (p *Proxy) proxy(w http.ResponseWriter, r *http.Request) {
	var responseHeader http.Header
	// If Sec-WebSocket-Protocol starts with "Bearer", respond in kind.
	// TODO(tmc): consider customizability/extension point here.
	if strings.HasPrefix(r.Header.Get("Sec-WebSocket-Protocol"), "Bearer") {
		responseHeader = http.Header{
			"Sec-WebSocket-Protocol": []string{"Bearer"},
		}
	}

	conn, err := upgrade.Upgrade(w, r, responseHeader)
	if err != nil {
		p.logger.Warnln("error upgrading websocket:", err)
		return
	}
	defer conn.Close()

	conn.SetReadLimit(maxMessageSize)
	if err := conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Err(err).Msg("failed to set read deadline")
	}
	conn.SetPingHandler(nil)
	conn.SetPongHandler(func(string) error {
		return conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	requestBodyR, requestBodyW := io.Pipe()
	p.logger.Warnln("backend service only supports POST requests")
	request, err := http.NewRequest(http.MethodPost, r.URL.String(), requestBodyR)
	if err != nil {
		p.logger.Warnln("error preparing request:", err)
		return
	}

	for k, v := range r.Header {
		for i := range v {
			request.Header.Add(k, v[i])
		}
	}

	if swsp := r.Header.Get("Sec-WebSocket-Protocol"); swsp != "" {
		request.Header.Set("Authorization", strings.Replace(swsp, "Bearer, ", "Bearer ", 1))
	}

	// If token cookie is present, populate Authorization header from the cookie instead.
	if cookie, err := r.Cookie(p.tokenCookieName); err == nil {
		request.Header.Set("Authorization", "Bearer "+cookie.Value)
	}

	if m := r.URL.Query().Get(p.methodOverrideParam); m != "" {
		request.Method = m
	}

	if p.requestMutator != nil {
		request = p.requestMutator(r, request)
	}

	responseBodyR, responseBodyW := io.Pipe()
	response := newInMemoryResponseWriter(responseBodyW)
	go func() {
		<-ctx.Done()
		p.logger.Debugln("closing websocket io pipes")
		requestBodyW.CloseWithError(io.EOF)
		responseBodyW.CloseWithError(io.EOF)
		response.closed <- true
	}()

	go func() {
		defer cancelFn()
		p.h.ServeHTTP(response, request)
	}()

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	defer func() {
		log.Info().Msg("close websocket ping")
	}()

	// read loop -- take messages from websocket and write to http request
	go func() {
		defer cancelFn()
		for {
			select {
			case <-ticker.C:
				conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					log.Err(err).Msg("failed to write ping message")
				}
			case <-ctx.Done():
				p.logger.Debugln("read loop done")
				return
			default:
				p.logger.Debugln("[read] reading from socket.")
				_, payload, err := conn.ReadMessage()
				if err != nil {
					if isClosedConnError(err) {
						p.logger.Debugln("[read] websocket closed:", err)
						return
					}
					p.logger.Warnln("error reading websocket message:", err)
					return
				}

				p.logger.Debugln("[read] read payload:", string(payload))
				p.logger.Debugln("[read] writing to requestBody:")
				n, err := requestBodyW.Write(payload)
				requestBodyW.Write([]byte("\n"))
				p.logger.Debugln("[read] wrote to requestBody", n)
				if err != nil {
					p.logger.Warnln("[read] error writing message to upstream http server:", err)
					return
				}
			}
		}
	}()

	// write loop -- take messages from response and write to websocket
	scanner := bufio.NewScanner(responseBodyR)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			p.logger.Warnln("[write] empty scan", scanner.Err())
			continue
		}

		p.logger.Debugln("[write] scanned", scanner.Text())
		if err = conn.WriteMessage(websocket.TextMessage, scanner.Bytes()); err != nil {
			p.logger.Warnln("[write] error writing websocket message:", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		p.logger.Warnln("scanner err:", err)
	}
}

type inMemoryResponseWriter struct {
	io.Writer
	header http.Header
	code   int
	closed chan bool
}

func newInMemoryResponseWriter(w io.Writer) *inMemoryResponseWriter {
	return &inMemoryResponseWriter{
		Writer: w,
		header: http.Header{},
		closed: make(chan bool, 1),
	}
}

func (w *inMemoryResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
func (w *inMemoryResponseWriter) Header() http.Header {
	return w.header
}
func (w *inMemoryResponseWriter) WriteHeader(code int) {
	w.code = code
}
func (w *inMemoryResponseWriter) CloseNotify() <-chan bool {
	return w.closed
}
func (w *inMemoryResponseWriter) Flush() {}
