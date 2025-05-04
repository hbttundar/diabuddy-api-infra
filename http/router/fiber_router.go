package router

import (
	"github.com/gofiber/fiber/v2"
)

type FiberRouter struct {
	engine *fiber.App
}

func NewFiberEngine(app *fiber.App) Engine {
	return &FiberRouter{engine: app}
}

func (f *FiberRouter) GET(path string, handler RouteHandler) {
	f.engine.Get(path, handler.(func(*fiber.Ctx) error))
}

func (f *FiberRouter) POST(path string, handler RouteHandler) {
	f.engine.Post(path, handler.(func(*fiber.Ctx) error))
}

func (f *FiberRouter) PUT(path string, handler RouteHandler) {
	f.engine.Put(path, handler.(func(*fiber.Ctx) error))
}

func (f *FiberRouter) PATCH(path string, handler RouteHandler) {
	f.engine.Patch(path, handler.(func(*fiber.Ctx) error))
}

func (f *FiberRouter) DELETE(path string, handler RouteHandler) {
	f.engine.Delete(path, handler.(func(*fiber.Ctx) error))
}

func (f *FiberRouter) Use(middleware ...Middleware) {
	for _, m := range middleware {
		f.engine.Use(m.(func(*fiber.Ctx) error))
	}
}

func (f *FiberRouter) Run(addr string) error {
	return f.engine.Listen(addr)
}
