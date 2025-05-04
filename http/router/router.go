package router

type RouteHandler interface{}
type Middleware interface{}

type Engine interface {
	GET(path string, handler RouteHandler)
	POST(path string, handler RouteHandler)
	PUT(path string, handler RouteHandler)
	PATCH(path string, handler RouteHandler)
	DELETE(path string, handler RouteHandler)
	Use(middleware ...Middleware)
	Run(addr string) error
}

type Option func(*Router)

type Router struct {
	engine     Engine
	middleware []Middleware
}

func NewRouter(opts ...Option) *Router {
	r := &Router{}
	for _, opt := range opts {
		opt(r)
	}
	return r.applyMiddleware()
}

func (r *Router) applyMiddleware() *Router {
	if r.engine != nil && len(r.middleware) > 0 {
		r.engine.Use(r.middleware...)
	}
	return r
}

func WithEngine(engine Engine) Option {
	return func(r *Router) {
		r.engine = engine
	}
}

func WithMiddleware(mw ...Middleware) Option {
	return func(r *Router) {
		r.middleware = append(r.middleware, mw...)
	}
}

func (r *Router) Engine() Engine {
	return r.engine
}

func (r *Router) GET(path string, handler RouteHandler) {
	if r.engine != nil {
		r.engine.GET(path, handler)
	}
}

func (r *Router) POST(path string, handler RouteHandler) {
	if r.engine != nil {
		r.engine.POST(path, handler)
	}
}

func (r *Router) PUT(path string, handler RouteHandler) {
	if r.engine != nil {
		r.engine.PUT(path, handler)
	}
}

func (r *Router) PATCH(path string, handler RouteHandler) {
	if r.engine != nil {
		r.engine.PATCH(path, handler)
	}
}

func (r *Router) DELETE(path string, handler RouteHandler) {
	if r.engine != nil {
		r.engine.DELETE(path, handler)
	}
}

func (r *Router) Use(mw ...Middleware) *Router {
	if r.engine != nil {
		r.engine.Use(mw...)
	}
	return r
}

func (r *Router) Run(addr string) error {
	if r.engine != nil {
		return r.engine.Run(addr)
	}
	return nil
}
