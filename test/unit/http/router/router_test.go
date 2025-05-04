package router_test

import (
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/hbttundar/diabuddy-api-infra/http/router"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGinRouterSatisfiesEngine(t *testing.T) {
	g := gin.Default()
	e := router.NewRouter(
		router.WithEngine(router.NewGinRouterWithEngine(g)),
	)

	assert.Implements(t, (*router.Engine)(nil), e.Engine())
}

func TestChiRouterSatisfiesEngine(t *testing.T) {
	c := chi.NewRouter()
	e := router.NewRouter(
		router.WithEngine(router.NewChiEngine(c)),
	)

	assert.Implements(t, (*router.Engine)(nil), e.Engine())
}

func TestFiberRouterSatisfiesEngine(t *testing.T) {
	f := fiber.New()
	e := router.NewRouter(
		router.WithEngine(router.NewFiberEngine(f)),
	)

	assert.Implements(t, (*router.Engine)(nil), e.Engine())
}

func TestHttpRouterSatisfiesEngine(t *testing.T) {
	h := httprouter.New()
	e := router.NewRouter(
		router.WithEngine(router.NewHttpRouterEngine(h)),
	)

	assert.Implements(t, (*router.Engine)(nil), e.Engine())
}

func TestRouterOptionsApply(t *testing.T) {
	g := gin.New()
	mw := gin.HandlerFunc(func(c *gin.Context) {
		c.Next()
	})

	r := router.NewRouter(
		router.WithEngine(router.NewGinRouterWithEngine(g)),
		router.WithMiddleware(mw),
	)

	assert.NotNil(t, r.Engine())
}

func TestGinRouterIntegration(t *testing.T) {
	g := gin.New()
	handlerCalled := false
	r := router.NewRouter(router.WithEngine(router.NewGinRouterWithEngine(g)))
	r.GET("/ping", gin.HandlerFunc(func(c *gin.Context) {
		handlerCalled = true
		c.String(http.StatusOK, "pong")
	}))

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
	assert.True(t, handlerCalled)
}

func TestHttpRouterIntegration(t *testing.T) {
	h := httprouter.New()
	handlerCalled := false

	e := router.NewRouter(router.WithEngine(router.NewHttpRouterEngine(h)))
	e.GET("/ping", httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handlerCalled = true
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	}))

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
	assert.True(t, handlerCalled)
}
