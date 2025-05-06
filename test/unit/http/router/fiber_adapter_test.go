package router_test

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/hbttundar/diabuddy-api-infra/http/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFiberRouter_GET(t *testing.T) {
	app := fiber.New()
	r := router.NewFiberAdapter(app)
	r.GET("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("test failed: %v", err)
	}
	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK || body.String() != "pong" {
		t.Errorf("unexpected response: %d %s", resp.StatusCode, body.String())
	}
}
