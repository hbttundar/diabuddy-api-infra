package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HttpRouter struct {
	engine *httprouter.Router
}

func NewHttpRouterEngine(r *httprouter.Router) Engine {
	return &HttpRouter{engine: r}
}

func (h *HttpRouter) GET(path string, handler RouteHandler) {
	h.engine.GET(path, handler.(httprouter.Handle))
}

func (h *HttpRouter) POST(path string, handler RouteHandler) {
	h.engine.POST(path, handler.(httprouter.Handle))
}

func (h *HttpRouter) PUT(path string, handler RouteHandler) {
	h.engine.PUT(path, handler.(httprouter.Handle))
}

func (h *HttpRouter) PATCH(path string, handler RouteHandler) {
	h.engine.PATCH(path, handler.(httprouter.Handle))
}

func (h *HttpRouter) DELETE(path string, handler RouteHandler) {
	h.engine.DELETE(path, handler.(httprouter.Handle))
}

func (h *HttpRouter) Use(middleware ...Middleware) {
	// Note: httprouter does not support native middleware,
	// you must wrap handlers manually before calling GET/POST/etc.
}

func (h *HttpRouter) Run(addr string) error {
	return http.ListenAndServe(addr, h.engine)
}
