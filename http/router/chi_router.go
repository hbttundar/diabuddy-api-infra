package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ChiRouter struct {
	engine *chi.Mux
}

func NewChiEngine(mux *chi.Mux) Engine {
	return &ChiRouter{engine: mux}
}

func (c *ChiRouter) GET(path string, handler RouteHandler) {
	c.engine.Method(http.MethodGet, path, handler.(http.Handler))
}

func (c *ChiRouter) POST(path string, handler RouteHandler) {
	c.engine.Method(http.MethodPost, path, handler.(http.Handler))
}

func (c *ChiRouter) PUT(path string, handler RouteHandler) {
	c.engine.Method(http.MethodPut, path, handler.(http.Handler))
}

func (c *ChiRouter) PATCH(path string, handler RouteHandler) {
	c.engine.Method(http.MethodPatch, path, handler.(http.Handler))
}

func (c *ChiRouter) DELETE(path string, handler RouteHandler) {
	c.engine.Method(http.MethodDelete, path, handler.(http.Handler))
}

func (c *ChiRouter) Use(middleware ...Middleware) {
	for _, m := range middleware {
		c.engine.Use(m.(func(http.Handler) http.Handler))
	}
}

func (c *ChiRouter) Run(addr string) error {
	return http.ListenAndServe(addr, c.engine)
}
