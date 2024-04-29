package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/pubgo/funk/log"
	"github.com/pubgo/funk/recovery"
	"github.com/pubgo/funk/result"
	"github.com/pubgo/funk/running"
	"github.com/pubgo/lava/core/logging/logkey"
	"github.com/rs/zerolog"
)

// New logger
func New(cfg *Config) log.Logger {
	defer recovery.Exit()

	level := zerolog.DebugLevel
	if cfg.Level != "" {
		level = result.Of(zerolog.ParseLevel(cfg.Level)).Expect("log level is invalid")
	}
	zerolog.SetGlobalLevel(level)

	logger := zerolog.New(os.Stdout).Level(level).With().Timestamp().Caller().Logger()
	if !cfg.AsJson {
		logger = logger.Output(&writer{
			ConsoleWriter: zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
				w.Out = os.Stdout
				w.TimeFormat = time.RFC3339
			}),
		})
	}

	// 全局log设置
	ee := logger.With().
		Str(logkey.Hostname, running.Hostname).
		Str(logkey.Project, running.Project).
		Str(logkey.Version, running.Version)

	if running.Namespace != "" {
		ee = ee.Str(logkey.Namespace, running.Namespace)
	}

	logger = ee.Logger()
	log.SetLogger(&logger)

	gl := log.New(&logger)
	for _, ext := range List() {
		ext(gl)
	}
	return gl
}

type writer struct {
	zerolog.ConsoleWriter
}

func (w writer) Write(p []byte) (n int, err error) {
	n, err = w.ConsoleWriter.Write(p)
	if err != nil {
		fmt.Println("invalid json: ", string(p))
	}
	return
}
