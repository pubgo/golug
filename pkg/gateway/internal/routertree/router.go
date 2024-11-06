package routertree

import (
	"fmt"
	"strings"
	"sync/atomic"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/pubgo/funk/assert"
	"github.com/pubgo/funk/errors"
	"github.com/pubgo/funk/generic"

	"github.com/pubgo/lava/pkg/gateway/internal/routerparser"
)

var (
	ErrPathNodeNotFound     = errors.New("path node not found")
	ErrNotFound             = errors.New("operation not found")
	ErrMethodNotAllowed     = errors.New("method not allowed")
	ErrVerbNotMatch         = errors.New("verb not match")
	ErrInvalidInput         = errors.New("invalid input")
	ErrRouterNotInitialized = errors.New("router not initialized")
)

// RouteTree represents a prefix tree for routing
type RouteTree struct {
	root  *node
	cache *lru.Cache[string, *MatchOperation]
	stats *routeStats
}

// routeStats 统计信息
type routeStats struct {
	matchCount  atomic.Int64
	cacheHits   atomic.Int64
	cacheMisses atomic.Int64
}

// node represents a node in the routing tree
type node struct {
	path     string
	children nodeChildren
	isWild   bool // whether this is a wildcard node (* or **)
	target   *routeTarget
}

// nodeChildren 优化子节点查找
type nodeChildren struct {
	static    map[string]*node
	wildcard  *node
	wildcard2 *node
}

func newNodeChildren() nodeChildren {
	return nodeChildren{
		static: make(map[string]*node),
	}
}

// routeTarget holds the endpoint information
type routeTarget struct {
	Method    string
	Path      string
	Operation string
	Verb      *string
	Vars      []*routerparser.PathVariable
	extras    map[string]any
}

// RouteOperation represents a route operation for external use
type RouteOperation struct {
	Method    string         `json:"method,omitempty"`
	Path      string         `json:"path,omitempty"`
	Operation string         `json:"operation,omitempty"`
	Verb      string         `json:"verb,omitempty"`
	Vars      []string       `json:"vars,omitempty"`
	Extras    map[string]any `json:"extras,omitempty"`
}

// MatchOperation represents a matched route operation
type MatchOperation struct {
	Method    string
	Path      string
	Operation string
	Verb      string
	Vars      []routerparser.PathFieldVar // reuse PathFieldVar from parser.go
	Extras    map[string]any
}

// NewRouteTree creates a new routing tree
func NewRouteTree() *RouteTree {
	cache := assert.Must1(lru.New[string, *MatchOperation](1000))
	return &RouteTree{
		root: &node{
			children: newNodeChildren(),
		},
		cache: cache,
		stats: &routeStats{},
	}
}

// Add adds a new route to the tree
func (r *RouteTree) Add(method string, path string, operation string, extras map[string]any) error {
	if r == nil || r.root == nil {
		return ErrRouterNotInitialized
	}

	pattern, err := routerparser.ParsePattern(path)
	if err != nil {
		return err
	}

	segments := make([]string, 0, len(pattern.Segments))
	for _, seg := range pattern.Segments {
		if seg != "" {
			segments = append(segments, seg)
		}
	}

	return r.addRoute(method, segments, pattern, operation, extras)
}

// addRoute 内部路由添加实现
func (r *RouteTree) addRoute(method string, segments []string, pattern *routerparser.Pattern, operation string, extras map[string]any) error {
	n := r.root
	methodNode := n.children.findChild(handlerMethod(method))
	if methodNode == nil {
		methodNode = &node{
			path:     handlerMethod(method),
			children: newNodeChildren(),
		}
		n.children.addChild(methodNode.path, methodNode)
	}

	n = methodNode
	for _, segment := range segments {
		child := n.children.findChild(segment)
		if child == nil {
			child = &node{
				path:     segment,
				children: newNodeChildren(),
				isWild:   segment == routerparser.Star || segment == routerparser.DoubleStar,
			}
			n.children.addChild(segment, child)
		}
		n = child
	}

	n.target = &routeTarget{
		Method:    method,
		Path:      pattern.Raw,
		Operation: operation,
		extras:    extras,
		Verb:      pattern.HttpVerb,
		Vars:      pattern.Variables,
	}

	return nil
}

// findChild finds a child node by path
func (n *nodeChildren) findChild(path string) *node {
	// 先查找静态路径
	if child, ok := n.static[path]; ok {
		return child
	}

	// 检查通配符
	if path == routerparser.Star {
		return n.wildcard
	}
	if path == routerparser.DoubleStar {
		return n.wildcard2
	}

	return nil
}

// addChild adds a child node
func (n *nodeChildren) addChild(path string, child *node) {
	if path == routerparser.Star {
		n.wildcard = child
	} else if path == routerparser.DoubleStar {
		n.wildcard2 = child
	} else {
		n.static[path] = child
	}
}

// Match finds a matching route for the given method and URL
func (r *RouteTree) Match(method, url string) (*MatchOperation, error) {
	if r == nil || r.root == nil {
		return nil, ErrRouterNotInitialized
	}

	r.stats.matchCount.Add(1)

	cacheKey := method + ":" + url
	if cached, ok := r.cache.Get(cacheKey); ok {
		r.stats.cacheHits.Add(1)
		return cached, nil
	}
	r.stats.cacheMisses.Add(1)

	result, err := r.match(method, url)
	if err == nil && result != nil {
		r.cache.Add(cacheKey, result)
	}
	return result, err
}

// List returns all routes in the tree
func (r *RouteTree) List() []RouteOperation {
	if r == nil || r.root == nil {
		return nil
	}

	var routes []RouteOperation
	var walk func(*node)

	walk = func(n *node) {
		if n == nil {
			return
		}

		// 如果节点有目标（即是一个路由终点），添加到结果中
		if n.target != nil {
			var vars []string
			// 只有当有变量时才初始化 vars
			if n.target.Vars != nil && len(n.target.Vars) > 0 {
				vars = make([]string, 0, len(n.target.Vars))
				for _, v := range n.target.Vars {
					vars = append(vars, strings.Join(v.FieldPath, "."))
				}
			}

			var verb string
			if n.target.Verb != nil {
				verb = *n.target.Verb
			}

			routes = append(routes, RouteOperation{
				Method:    n.target.Method,
				Path:      n.target.Path,
				Operation: n.target.Operation,
				Verb:      verb,
				Vars:      vars, // 如果没有变量，将保持为 nil
				Extras:    n.target.extras,
			})
		}

		// 遍历所有子节点
		for _, child := range n.children.static {
			walk(child)
		}
		if n.children.wildcard != nil {
			walk(n.children.wildcard)
		}
		if n.children.wildcard2 != nil {
			walk(n.children.wildcard2)
		}
	}

	// 遍历方法节点
	for _, methodNode := range r.root.children.static {
		walk(methodNode)
	}

	return routes
}

// walkNode recursively walks the route tree and collects all operations
func (r *RouteTree) walkNode(n *node, prefix string, ops *[]RouteOperation) {
	if n == nil {
		return
	}

	// 构建当前路径
	path := prefix
	if len(n.path) > 0 {
		if path == "" {
			path = n.path
		} else {
			path += "/" + n.path
		}
	}

	// 如果当前节点有目标操作，添加到结果中
	if n.target != nil {
		vars := make([]string, 0, len(n.target.Vars))
		for _, v := range n.target.Vars {
			vars = append(vars, strings.Join(v.FieldPath, "."))
		}

		*ops = append(*ops, RouteOperation{
			Method:    n.target.Method,
			Path:      n.target.Path,
			Operation: n.target.Operation,
			Verb:      generic.FromPtr(n.target.Verb),
			Vars:      vars,
			Extras:    n.target.extras,
		})
	}

	// 递归遍历所有子节点
	for _, child := range n.children.static {
		r.walkNode(child, path, ops)
	}
}

// Stats returns the router's statistics
func (r *RouteTree) Stats() map[string]interface{} {
	return map[string]interface{}{
		"total_matches": r.stats.matchCount.Load(),
		"cache_hits":    r.stats.cacheHits.Load(),
		"cache_misses":  r.stats.cacheMisses.Load(),
	}
}

// Helper functions
func handlerMethod(method string) string {
	return fmt.Sprintf("__%s__", strings.ToUpper(method))
}

// match finds a matching route
func (r *RouteTree) match(method, url string) (*MatchOperation, error) {
	// 解析URL
	urlInfo := parseURL(url)

	// 查找方法节点
	methodNode := r.root.children.findChild(handlerMethod(method))
	if methodNode == nil {
		return nil, ErrMethodNotAllowed
	}

	// 执行路径匹配
	result, err := r.matchPath(methodNode, urlInfo)
	if err != nil {
		return nil, err
	}

	// 验证动词
	if result.target.Verb != nil {
		if urlInfo.verb == "" || *result.target.Verb != urlInfo.verb {
			return nil, ErrVerbNotMatch
		}
	}

	return r.buildMatchOperation(result, urlInfo)
}

type urlInfo struct {
	segments []string
	verb     string
	raw      string
}

func parseURL(url string) *urlInfo {
	info := &urlInfo{raw: url}

	// 解析路径段
	paths := strings.Split(strings.Trim(url, "/"), "/")
	info.segments = make([]string, 0, len(paths))
	for _, p := range paths {
		if p != "" {
			info.segments = append(info.segments, p)
		}
	}

	// 解析动词
	if len(info.segments) > 0 {
		lastPath := info.segments[len(info.segments)-1]
		if idx := strings.LastIndex(lastPath, ":"); idx >= 0 {
			info.verb = lastPath[idx+1:]
			info.segments[len(info.segments)-1] = lastPath[:idx]
		}
	}

	return info
}

// matchPath 在给定的节点下匹配路径
func (r *RouteTree) matchPath(n *node, info *urlInfo) (*node, error) {
	if n == nil || info == nil {
		return nil, ErrInvalidInput
	}

	type matchState struct {
		node     *node
		segIndex int
		priority int // 添加优先级字段，数字越小优先级越高
	}

	var stack []matchState
	stack = append(stack, matchState{node: n, segIndex: 0, priority: 0})

	var bestMatch *node
	var bestPriority int = 1000 // 设置一个较大的初始值

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 如果找到完整匹配，并且优先级更高，则更新最佳匹配
		if current.segIndex == len(info.segments) {
			if current.node.target != nil && current.priority < bestPriority {
				bestMatch = current.node
				bestPriority = current.priority
			}
			continue
		}

		if current.segIndex >= len(info.segments) {
			continue
		}

		segment := info.segments[current.segIndex]
		children := &current.node.children

		// 按优先级顺序添加到栈中（后添加的先处理）

		// 3. ** 通配符 (最低优先级)
		if children.wildcard2 != nil {
			stack = append(stack, matchState{
				node:     children.wildcard2,
				segIndex: len(info.segments),
				priority: current.priority + 3,
			})
		}

		// 2. * 通配符 (中等优先级)
		if children.wildcard != nil {
			stack = append(stack, matchState{
				node:     children.wildcard,
				segIndex: current.segIndex + 1,
				priority: current.priority + 2,
			})
		}

		// 1. 静态路径 (最高优先级)
		if child, ok := children.static[segment]; ok && child != nil {
			stack = append(stack, matchState{
				node:     child,
				segIndex: current.segIndex + 1,
				priority: current.priority + 1,
			})
		}
	}

	if bestMatch != nil {
		return bestMatch, nil
	}

	return nil, ErrPathNodeNotFound
}

// buildMatchOperation 构建匹配结果
func (r *RouteTree) buildMatchOperation(n *node, info *urlInfo) (*MatchOperation, error) {
	if n == nil || n.target == nil || info == nil {
		return nil, ErrNotFound
	}

	pattern := &routerparser.Pattern{
		Raw:       n.target.Path,
		HttpVerb:  n.target.Verb,
		Variables: n.target.Vars,
	}

	var vars []routerparser.PathFieldVar
	var err error

	// 只有在有变量定义时才进行变量匹配
	if len(pattern.Variables) > 0 {
		vars, err = pattern.Match(info.segments, info.verb)
		if err != nil {
			return nil, err
		}
	}

	return &MatchOperation{
		Method:    n.target.Method,
		Path:      n.target.Path,
		Operation: n.target.Operation,
		Verb:      info.verb,
		Vars:      vars,
		Extras:    n.target.extras,
	}, nil
}
