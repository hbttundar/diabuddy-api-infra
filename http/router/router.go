// router package: universal router abstraction that supports all engines
package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
)

// EngineType is the enum type for supported engines
type EngineType string

const (
	GinEngine   EngineType = "gin"
	EchoEngine  EngineType = "echo"
	FiberEngine EngineType = "fiber"
	ChiEngine   EngineType = "chi"
)

// Middleware defines the generic middleware type.
type Middleware interface{}

// Adapter represents the routing adapter with optional capabilities.
type Adapter interface {
	GET(path string, handlers ...interface{})
	POST(path string, handlers ...interface{})
	PUT(path string, handlers ...interface{})
	PATCH(path string, handlers ...interface{})
	DELETE(path string, handlers ...interface{})
	ServeHTTP(w http.ResponseWriter, req *http.Request)

	Adapter() Adapter
}

// RouterOption configures a Router during setup.
type RouterOption func(*Router)

// Router uses an Adapter to route requests.
type Router struct {
	adapter    Adapter
	middleware []Middleware
}

// NewRouterFromType selects adapter by engine type and returns a ready router
func NewRouterFromType(engineType EngineType, middleware ...Middleware) *Router {
	var adapter Adapter

	switch strings.ToLower(string(engineType)) {
	case string(GinEngine):
		adapter = NewGinAdapter(gin.New())
	case string(EchoEngine):
		adapter = NewEchoAdapter(echo.New())
	case string(FiberEngine):
		adapter = NewFiberAdapter(fiber.New())
	case string(ChiEngine):
		adapter = NewChiAdapter(chi.NewRouter())
	default:
		panic("unsupported engine type: " + string(engineType))
	}

	r := &Router{adapter: adapter, middleware: middleware}
	r.applyMiddleware()
	return r
}

func (r *Router) applyMiddleware() {
	if mwCapable, ok := r.adapter.(interface{ Use(...Middleware) }); ok && len(r.middleware) > 0 {
		mwCapable.Use(r.middleware...)
	}
}

func (r *Router) GET(path string, handlers ...interface{}) {
	r.adapter.GET(path, handlers...)
}
func (r *Router) POST(path string, handlers ...interface{}) {
	r.adapter.POST(path, handlers...)
}
func (r *Router) PUT(path string, handlers ...interface{}) {
	r.adapter.PUT(path, handlers...)
}
func (r *Router) PATCH(path string, handlers ...interface{}) {
	r.adapter.PATCH(path, handlers...)
}
func (r *Router) DELETE(path string, handlers ...interface{}) {
	r.adapter.DELETE(path, handlers...)
}

func (r *Router) Adapter() Adapter {
	return r.adapter
}

func (r *Router) Use(mw ...Middleware) *Router {
	if mwCapable, ok := r.adapter.(interface{ Use(...Middleware) }); ok {
		mwCapable.Use(mw...)
	}
	return r
}

func (r *Router) Run(addr ...string) error {
	if runner, ok := r.adapter.(interface{ Run(...string) error }); ok {
		return runner.Run(addr...)
	}
	return fmt.Errorf("adapter does not implement Run")
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if h, ok := r.adapter.(http.Handler); ok {
		h.ServeHTTP(w, req)
		return
	}
	panic(fmt.Sprintf("router.adapter %T does not implement http.Handler", r.adapter))
}
