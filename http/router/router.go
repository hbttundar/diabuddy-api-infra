package router

// RouteHandler defines the signature for handling a route.
type RouteHandler interface{}

// Middleware wraps a RouteHandler, allowing pre- and post-processing.
type Middleware interface{}

// Engine defines the minimal interface any HTTP engine must implement.
type Engine interface {
	GET(path string, handler RouteHandler)
	POST(path string, handler RouteHandler)
	PUT(path string, handler RouteHandler)
	PATCH(path string, handler RouteHandler)
	DELETE(path string, handler RouteHandler)
}

// HasMiddleware marks engines that support middleware registration.
type HasMiddleware interface {
	// Use registers one or more middleware handlers.
	Use(middleware ...Middleware)
}

// Runnable marks engines that support running the server.
type Runnable interface {
	// Run starts the server with the provided start function.
	Run(start func() error) error
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

// Run starts the underlying engine using the provided start function.
func (r *Router) Run(start func() error) error {
	if runnable, ok := r.engine.(Runnable); ok {
		return runnable.Run(start)
	}
	return start()
}
