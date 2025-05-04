# diabuddy-api-infra

Shared infrastructure module for the DiaBuddy microservices platform.
This package provides reusable, domain-agnostic building blocks to standardize:

* Database abstraction and lifecycle management
* Common repository patterns
* HTTP router and client setup
* Secure & general-purpose helper utilities

---

## Modules

### `database/`

* `Connection` interface: abstract DB connection lifecycle
* `PostgresConnection`: PostgreSQL implementation
* `WrapInTransaction()`: transaction + advisory lock helper
* Constants: `InsertOperation`, `UpdateOperation`, etc.

### `persistence/`

* `BaseRepository`: common DB query/scan/exec logic
* Shared pagination support (`paginator`, `Pagination`)

### `http/`

#### `router/` – Pluggable HTTP Router Engine

* Abstract `Router` system with functional options
* `Engine` interface defines `GET`, `POST`, `Use`, `Run`, etc.
* Ready-to-use adapters:

    * `GinRouter` (`gin-gonic/gin`)
    * `ChiRouter` (`go-chi/chi`)
    * `FiberRouter` (`gofiber/fiber`)
    * `HttpRouter` (`julienschmidt/httprouter`)
* Supports middleware injection via `WithMiddleware` excepts `HttpRouter` that didn't support middleware

#### `client/`

* `DefaultHTTPClient()`: shared, tuned `http.Client` for external API calls

### `helpers/`

* `hasher`: bcrypt-based password hash/verify
* `datetime`: UTC helpers for consistent time handling
* `utils`: generic helpers (`ToPointer`, `Coalesce`, etc.)

---

## Usage

### Setup a Router in Your Microservice

```go
import (
  "github.com/hbttundar/diabuddy-api-infra/http/router"
  "github.com/gin-gonic/gin"
)

func NewRouter() *router.Router {
  return router.NewRouter(
    router.WithEngine(router.NewGinRouter()),
    router.WithMiddleware(gin.Logger(), gin.Recovery()),
  )
}
```

### Start the Engine

```go
r := NewRouter()
_ = r.Run(":8080")
```

Licensed under MIT.
