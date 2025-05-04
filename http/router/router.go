package router

type Router struct {
	engine     any
	middleware []any // let each consumer cast properly
}

type Option func(*Router)

func NewRouter(opts ...Option) *Router {
	r := &Router{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func WithEngine(engine any) Option {
	return func(r *Router) {
		r.engine = engine
	}
}

func WithMiddleware(mw ...any) Option {
	return func(r *Router) {
		r.middleware = append(r.middleware, mw...)
	}
}

func (r *Router) Engine() any {
	return r.engine
}

func (r *Router) Middleware() []any {
	return r.middleware
}
