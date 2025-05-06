package router_test

import (
	"github.com/hbttundar/diabuddy-api-infra/http/router"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoRouter_GET(t *testing.T) {
	e := echo.New()
	r := router.NewEchoAdapter(e)
	r.GET("/ping", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	}))

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK || rec.Body.String() != "pong" {
		t.Errorf("unexpected response: %d %s", rec.Code, rec.Body.String())
	}
}
