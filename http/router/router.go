package router

type RouteHandler interface{}
type Middleware interface{}

type Engine interface {
	GET(path string, handler RouteHandler)
	POST(path string, handler RouteHandler)
	PUT(path string, handler RouteHandler)
	PATCH(path string, handler RouteHandler)
	DELETE(path string, handler RouteHandler)
	Run(addr string) error
}

// HasMiddleware marks engines that support middleware registration.
type HasMiddleware interface {
	// Use registers one or more middleware handlers.
	Use(middleware ...Middleware)
}

// RouterOption configures a Router during creation.
type RouterOption func(*Router)

// Router composes an engine and optional middleware, applying them on startup.
type Router struct {
	engine     Engine
	middleware []Middleware
}

// NewRouter constructs a Router, applies all options, and registers middleware if supported.
func NewRouter(opts ...RouterOption) *Router {
	r := &Router{}
	for _, opt := range opts {
		opt(r)
	}
	// Apply middleware if engine supports it
	if mwEngine, ok := r.engine.(HasMiddleware); ok && len(r.middleware) > 0 {
		mwEngine.Use(r.middleware...)
	}
	return r
}

// WithEngine sets the underlying Engine.
func WithEngine(e Engine) RouterOption {
	return func(r *Router) {
		r.engine = e
	}
}

// WithMiddleware appends middleware to the Router before initialization.
func WithMiddleware(mw ...Middleware) RouterOption {
	return func(r *Router) {
		r.middleware = append(r.middleware, mw...)
	}
}

// Engine exposes the underlying Engine.
func (r *Router) Engine() Engine {
	return r.engine
}

// GET registers a GET route on the engine.
func (r *Router) GET(path string, handler RouteHandler) {
	if r.engine != nil {
		r.engine.GET(path, handler)
	}
}

// POST registers a POST route on the engine.
func (r *Router) POST(path string, handler RouteHandler) {
	if r.engine != nil {
		r.engine.POST(path, handler)
	}
}

// PUT registers a PUT route on the engine.
func (r *Router) PUT(path string, handler RouteHandler) {
	if r.engine != nil {
		r.engine.PUT(path, handler)
	}
}

// PATCH registers a PATCH route on the engine.
func (r *Router) PATCH(path string, handler RouteHandler) {
	if r.engine != nil {
		r.engine.PATCH(path, handler)
	}
}

// DELETE registers a DELETE route on the engine.
func (r *Router) DELETE(path string, handler RouteHandler) {
	if r.engine != nil {
		r.engine.DELETE(path, handler)
	}
}

// Use registers additional middleware on-the-fly if the engine supports it.
func (r *Router) Use(mw ...Middleware) *Router {
	if mwEngine, ok := r.engine.(HasMiddleware); ok {
		mwEngine.Use(mw...)
	}
	return r
}

// Run starts the underlying engine on the given address.
func (r *Router) Run(addr string) error {
	if r.engine != nil {
		return r.engine.Run(addr)
	}
	return nil
}
