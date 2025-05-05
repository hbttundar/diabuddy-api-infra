package router_test

import (
	"github.com/go-chi/chi/v5"
	"github.com/hbttundar/diabuddy-api-infra/http/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChiRouter_GET(t *testing.T) {
	r := chi.NewRouter()
	a := router.NewChiAdapter(r)
	a.GET("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}))

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK || rec.Body.String() != "pong" {
		t.Errorf("unexpected response: %d %s", rec.Code, rec.Body.String())
	}
}
