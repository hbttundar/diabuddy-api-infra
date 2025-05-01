package http

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "net/http"
	"time"
)

// SetupRouter initializes a Gin engine with default middleware and optional route setup.
func SetupRouter(registerRoutes func(r *gin.Engine)) *gin.Engine {
	r := gin.New()

	// Apply standard middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(timeoutMiddleware(10 * time.Second))

	// Let the caller define routes
	registerRoutes(r)

	return r
}

// timeoutMiddleware limits the lifetime of requests.
func timeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
