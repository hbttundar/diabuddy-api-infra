package router

import "github.com/gin-gonic/gin"

type GinRouter struct {
	engine *gin.Engine
}

func NewGinRouter() *GinRouter {
	return &GinRouter{engine: gin.Default()}
}

func NewGinRouterWithEngine(e *gin.Engine) *GinRouter {
	return &GinRouter{engine: e}
}

func (g *GinRouter) GET(path string, handler RouteHandler) {
	g.engine.GET(path, handler.(gin.HandlerFunc))
}

func (g *GinRouter) POST(path string, handler RouteHandler) {
	g.engine.POST(path, handler.(gin.HandlerFunc))
}

func (g *GinRouter) PUT(path string, handler RouteHandler) {
	g.engine.PUT(path, handler.(gin.HandlerFunc))
}

func (g *GinRouter) PATCH(path string, handler RouteHandler) {
	g.engine.PATCH(path, handler.(gin.HandlerFunc))
}

func (g *GinRouter) DELETE(path string, handler RouteHandler) {
	g.engine.DELETE(path, handler.(gin.HandlerFunc))
}

func (g *GinRouter) Use(middleware ...Middleware) {
	for _, m := range middleware {
		g.engine.Use(m.(gin.HandlerFunc))
	}
}

func (g *GinRouter) Run(addr string) error {
	return g.engine.Run(addr)
}
