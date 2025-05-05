package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type FiberAdapter struct {
	engine *fiber.App
}

func NewFiberAdapter(engin *fiber.App) *FiberAdapter {
	return &FiberAdapter{engine: engin}
}

func (adapter *FiberAdapter) GET(path string, handlers ...interface{}) {
	adapter.engine.Get(path, adapter.resolveFiberHandlers(handlers...)...)
}
func (adapter *FiberAdapter) POST(path string, handlers ...interface{}) {
	adapter.engine.Post(path, adapter.resolveFiberHandlers(handlers...)...)
}
func (adapter *FiberAdapter) PUT(path string, handlers ...interface{}) {
	adapter.engine.Put(path, adapter.resolveFiberHandlers(handlers...)...)
}
func (adapter *FiberAdapter) PATCH(path string, handlers ...interface{}) {
	adapter.engine.Patch(path, adapter.resolveFiberHandlers(handlers...)...)
}
func (adapter *FiberAdapter) DELETE(path string, handlers ...interface{}) {
	adapter.engine.Delete(path, adapter.resolveFiberHandlers(handlers...)...)
}

func (adapter *FiberAdapter) Adapter() Adapter {
	return adapter
}

func (adapter *FiberAdapter) FiberEngine() *fiber.App {
	return adapter.engine
}

func (adapter *FiberAdapter) Run(addr ...string) error {
	host := ":8080"
	if len(addr) > 0 {
		host = addr[0]
	}
	return adapter.engine.Listen(host)
}

func (adapter *FiberAdapter) Use(mw ...Middleware) {
	for _, m := range mw {
		if fn, ok := m.(fiber.Handler); ok {
			adapter.engine.Use(fn)
		} else {
			panic("FiberAdapter middleware must be fiber.Handler")
		}
	}
}

func (adapter *FiberAdapter) resolveFiberHandlers(handlers ...interface{}) []fiber.Handler {
	result := make([]fiber.Handler, len(handlers))
	for i, h := range handlers {
		fn, ok := h.(fiber.Handler)
		if !ok {
			panic(fmt.Sprintf("handler at index %d is not a fiber.Handler", i))
		}
		result[i] = fn
	}
	return result
}
