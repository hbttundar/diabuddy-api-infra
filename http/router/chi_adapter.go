package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ChiAdapter struct {
	router chi.Router
}

func NewChiAdapter(r chi.Router) *ChiAdapter {
	return &ChiAdapter{router: r}
}

func (adapter *ChiAdapter) GET(path string, handlers ...interface{}) {
	adapter.router.Method(http.MethodGet, path, adapter.resolveChiHandler(handlers...))
}
func (adapter *ChiAdapter) POST(path string, handlers ...interface{}) {
	adapter.router.Method(http.MethodPost, path, adapter.resolveChiHandler(handlers...))
}
func (adapter *ChiAdapter) PUT(path string, handlers ...interface{}) {
	adapter.router.Method(http.MethodPut, path, adapter.resolveChiHandler(handlers...))
}
func (adapter *ChiAdapter) PATCH(path string, handlers ...interface{}) {
	adapter.router.Method(http.MethodPatch, path, adapter.resolveChiHandler(handlers...))
}
func (adapter *ChiAdapter) DELETE(path string, handlers ...interface{}) {
	adapter.router.Method(http.MethodDelete, path, adapter.resolveChiHandler(handlers...))
}

func (adapter *ChiAdapter) Adapter() Adapter {
	return adapter
}

func (adapter *ChiAdapter) ChiEngine() chi.Router {
	return adapter.router
}

func (adapter *ChiAdapter) Run(addr ...string) error {
	host := ":8080"
	if len(addr) > 0 {
		host = addr[0]
	}
	return http.ListenAndServe(host, adapter.router)
}

func (adapter *ChiAdapter) Use(middlewares ...Middleware) {
	for _, m := range middlewares {
		if fn, ok := m.(func(http.Handler) http.Handler); ok {
			adapter.router.Use(fn)
		} else {
			panic("ChiAdapter middleware must be func(http.Handler) http.Handler")
		}
	}
}

func (adapter *ChiAdapter) resolveChiHandler(handlers ...interface{}) http.Handler {
	if len(handlers) != 1 {
		panic("ChiAdapter expects exactly one handler")
	}
	h := handlers[0]
	switch fn := h.(type) {
	case http.Handler:
		return fn
	case http.HandlerFunc:
		return fn
	case func(http.ResponseWriter, *http.Request):
		return http.HandlerFunc(fn)
	default:
		panic(fmt.Sprintf("chi adapter: handler is %T, want http.Handler or http.HandlerFunc", h))
	}
}

func (adapter *ChiAdapter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	adapter.router.ServeHTTP(w, req)
}
