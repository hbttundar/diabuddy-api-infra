package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type EchoAdapter struct {
	engine *echo.Echo
}

func NewEchoAdapter(engin *echo.Echo) *EchoAdapter {
	return &EchoAdapter{engine: engin}
}

func (adapter *EchoAdapter) GET(path string, handlers ...interface{}) {
	adapter.engine.GET(path, adapter.resolveEchoHandler(handlers...))
}
func (adapter *EchoAdapter) POST(path string, handlers ...interface{}) {
	adapter.engine.POST(path, adapter.resolveEchoHandler(handlers...))
}
func (adapter *EchoAdapter) PUT(path string, handlers ...interface{}) {
	adapter.engine.PUT(path, adapter.resolveEchoHandler(handlers...))
}
func (adapter *EchoAdapter) PATCH(path string, handlers ...interface{}) {
	adapter.engine.PATCH(path, adapter.resolveEchoHandler(handlers...))
}
func (adapter *EchoAdapter) DELETE(path string, handlers ...interface{}) {
	adapter.engine.DELETE(path, adapter.resolveEchoHandler(handlers...))
}

func (adapter *EchoAdapter) Adapter() Adapter {
	return adapter
}

func (adapter *EchoAdapter) EchoEngine() *echo.Echo {
	return adapter.engine
}

func (adapter *EchoAdapter) Run(addr ...string) error {
	host := ":8080"
	if len(addr) > 0 {
		host = addr[0]
	}
	return adapter.engine.Start(host)
}

func (adapter *EchoAdapter) Use(middlewares ...Middleware) {
	for _, m := range middlewares {
		if fn, ok := m.(echo.MiddlewareFunc); ok {
			adapter.engine.Use(fn)
		} else {
			panic("EchoAdapter middleware must be echo.MiddlewareFunc")
		}
	}
}

func (adapter *EchoAdapter) resolveEchoHandler(handlers ...interface{}) echo.HandlerFunc {
	if len(handlers) != 1 {
		panic("EchoAdapter expects exactly one handler")
	}
	switch fn := handlers[0].(type) {
	case echo.HandlerFunc:
		return fn
	case func(echo.Context) error:
		return echo.HandlerFunc(fn)
	default:
		panic("EchoAdapter expects exactly one handler")
	}
}

func (adapter *EchoAdapter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	adapter.engine.ServeHTTP(w, req)
}
