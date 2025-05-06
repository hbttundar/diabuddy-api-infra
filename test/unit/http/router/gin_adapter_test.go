package router_test

import (
	"github.com/gin-gonic/gin"
	"github.com/hbttundar/diabuddy-api-infra/http/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGinRouter_GET(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.NewRouterFromType(router.GinEngine)
	r.GET("/ping", gin.HandlerFunc(func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	r.Adapter().(*router.GinAdapter).GinEngine().ServeHTTP(w, req)

	if w.Code != http.StatusOK || w.Body.String() != "pong" {
		t.Errorf("unexpected response: %d %s", w.Code, w.Body.String())
	}
}
