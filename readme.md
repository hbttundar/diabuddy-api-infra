# diabuddy-api-infra

Shared infrastructure module for the DiaBuddy microservices platform. This package provides reusable, low-level building blocks for database, HTTP, time, and general-purpose operations.

---

## 📦 Features

### Database
- `database.Connection`: Unified DB interface
- `NewPostgresConnection`: PG implementation
- Constants for CRUD operations (e.g. `InsertOperation`, `UpdateOperation`)

### HTTP
- `SetupGinRouter`: Bootstraps Gin with middleware and custom routes
- `DefaultHTTPClient`: Sensible HTTP client for service-to-service calls

### Helpers
- `password`: bcrypt hashing and comparison helpers
- `datetime`: date formatting, parsing, UTC utils
- `utils`: generic helpers like `ToPointer`, `Coalesce`

---

## 🔗 Usage

### In Your Microservice
```go
import "github.com/hbttundar/diabuddy-api-infra/database"

conn := database.NewPostgresConnection(...)
```

### In Testkit
```go
import "github.com/hbttundar/diabuddy-api-infra/helpers"

now := helpers.NowUTC()
ptr := helpers.ToPointer("value")
```

---

## 🧱 Structure

```
database/     # Shared DB interface + Postgres impl + constants
http/         # Shared Gin router + HTTP client
helpers/      # General-purpose helpers (password, datetime, etc.)
```

---

## ✅ Best Practices
- Keep shared domain-agnostic code here
- Do NOT place domain-specific logic (e.g. User, Profile) in infra
- Keep this library focused, reusable, and stable

---

## 🛠 Future Additions
- Kafka and Elasticsearch client setups
- Retry and circuit breaker utilities
- Service discovery hooks (Consul, etc.)

---

Licensed under MIT.
