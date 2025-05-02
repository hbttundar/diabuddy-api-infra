# diabuddy-api-infra

Shared infrastructure module for the DiaBuddy microservices platform. This package provides reusable, domain-agnostic building blocks for:
- Database abstraction and lifecycle management
- Common repository operations
- HTTP client and router setup
- General-purpose and secure helpers (datetime, password)

---

## 📦 Modules

### 🗃️ database/
- `Connection` interface (abstract DB contract)
- `PostgresConnection` implementation
- `WrapInTransaction()` with advisory lock
- Constants for CRUD operations (e.g. `InsertOperation`, `UpdateOperation`)

### 🗄️ persistence/
- `BaseRepository`: common SQL ops (Exec, ScanRow, ParseResult)

### 🌐 http/
- `SetupGinRouter`: initializes Gin engine with shared middleware
- `DefaultHTTPClient`: opinionated shared `http.Client`

### 🔐 helpers/
- `hasher`: bcrypt hash & verify password
- `datetime`: UTC utils, format/parse helpers
- `utils`: general helpers (`ToPointer`, `Coalesce`, etc.)

---

## 🔗 Usage

### From a microservice
```go
import "github.com/hbttundar/diabuddy-api-infra/database"
import "github.com/hbttundar/diabuddy-api-infra/persistence"

conn := database.NewPostgresConnection(...) // or use default config
repo := persistence.NewBaseRepository(conn)
```

### From testkit
```go
import "github.com/hbttundar/diabuddy-api-infra/helpers"

now := helpers.NowUTC()
hash, _ := helpers.HashPassword("secret")
```

---

## 🧱 Directory Structure

```
database/     # Abstract DB interfaces and PG implementation
persistence/  # Base repository abstraction for microservices
http/         # Gin router + HTTP client
helpers/      # Shared logic for time, passwords, general use
```

---

## ✅ Best Practices
- This module should contain **only infrastructure** logic
- Do not mix domain-specific models or services
- Use `diabuddy-errors` for consistent API error handling
- Use `diabuddy-api-config` for all env and DB config needs

---

## 🔧 Future Enhancements
- Kafka/Elasticsearch clients
- Circuit breaker, retry helpers
- Observability (OpenTelemetry, metrics)

---

Licensed under MIT
