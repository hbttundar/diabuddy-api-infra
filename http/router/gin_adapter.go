package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type GinAdapter struct {
	engine *gin.Engine
}

func NewGinAdapter(engin *gin.Engine) *GinAdapter {
	return &GinAdapter{engine: engin}
}

func (adapter *GinAdapter) GET(path string, handlers ...interface{}) {
	h := adapter.resolveGinHandlers(handlers...)
	adapter.engine.GET(path, h...)
}
func (adapter *GinAdapter) POST(path string, handlers ...interface{}) {
	h := adapter.resolveGinHandlers(handlers...)
	adapter.engine.POST(path, h...)
}
func (adapter *GinAdapter) PUT(path string, handlers ...interface{}) {
	h := adapter.resolveGinHandlers(handlers...)
	adapter.engine.PUT(path, h...)
}
func (adapter *GinAdapter) PATCH(path string, handlers ...interface{}) {
	h := adapter.resolveGinHandlers(handlers...)
	adapter.engine.PATCH(path, h...)
}
func (adapter *GinAdapter) DELETE(path string, handlers ...interface{}) {
	h := adapter.resolveGinHandlers(handlers...)
	adapter.engine.DELETE(path, h...)
}

func (adapter *GinAdapter) Adapter() Adapter {
	return adapter
}

func (adapter *GinAdapter) GinEngine() *gin.Engine {
	return adapter.engine
}

func (adapter *GinAdapter) Run(addr ...string) error {
	host := ":8080"
	if len(addr) > 0 {
		host = addr[0]
	}
	return adapter.engine.Run(host)
}

func (adapter *GinAdapter) Use(mw ...Middleware) {
	for _, m := range mw {
		if fn, ok := m.(gin.HandlerFunc); ok {
			adapter.engine.Use(fn)
		}
	}
}
func (adapter *GinAdapter) resolveGinHandlers(handlers ...interface{}) []gin.HandlerFunc {
	result := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		fn, ok := h.(gin.HandlerFunc)
		if !ok {
			panic(fmt.Sprintf("handler at index %d is not a gin.HandlerFunc", i))
		}
		result[i] = fn
	}
	return result
}
