package router

import (
	"github.com/labstack/echo/v4"
)

type EchoRouter struct {
	engine *echo.Echo
}

func NewEchoEngine(e *echo.Echo) Engine {
	return &EchoRouter{engine: e}
}

func (e *EchoRouter) GET(path string, handler RouteHandler) {
	e.engine.GET(path, handler.(echo.HandlerFunc))
}

func (e *EchoRouter) POST(path string, handler RouteHandler) {
	e.engine.POST(path, handler.(echo.HandlerFunc))
}

func (e *EchoRouter) PUT(path string, handler RouteHandler) {
	e.engine.PUT(path, handler.(echo.HandlerFunc))
}

func (e *EchoRouter) PATCH(path string, handler RouteHandler) {
	e.engine.PATCH(path, handler.(echo.HandlerFunc))
}

func (e *EchoRouter) DELETE(path string, handler RouteHandler) {
	e.engine.DELETE(path, handler.(echo.HandlerFunc))
}

func (e *EchoRouter) Use(middleware ...Middleware) {
	for _, m := range middleware {
		e.engine.Use(m.(echo.MiddlewareFunc))
	}
}

func (e *EchoRouter) Run(addr string) error {
	return e.engine.Start(addr)
}
