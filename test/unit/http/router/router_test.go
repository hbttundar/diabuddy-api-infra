package router_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/hbttundar/diabuddy-api-infra/http/router"
)

func TestRouter_WithEngineAndMiddleware(t *testing.T) {
	engine := gin.New()
	middleware1 := gin.Logger()
	middleware2 := gin.Recovery()

	r := router.NewRouter(
		router.WithEngine(engine),
		router.WithMiddleware(middleware1, middleware2),
	)

	assert.Equal(t, engine, r.Engine())
	assert.Len(t, r.Middleware(), 2)
}
