# REED Backend (chi-based scaffold)

This folder contains a minimal scaffold for a Go backend using the chi router.
The implementation is intentionally left mostly empty — this layout is meant to make responsibilities and file locations clear so you can implement functionality iteratively.

Directory layout (important files):

- `cmd/server/main.go` — application entrypoint (placeholder)
- `internal/server` — server/router setup (contains a `NewRouter` function)
- `internal/routes` — centralized route wiring
- `internal/controllers` — HTTP controllers (e.g. `ModelController`)
- `internal/models` — domain model definitions
- `internal/repository` — repository interfaces and implementations
- `internal/services` — business logic and service interfaces
- `internal/middleware` — HTTP middleware
- `config` — configuration loader and config structs

Next steps:

1. Set the module path in `go.mod` (if not already set) and add dependencies (e.g. `github.com/go-chi/chi/v5`).
2. Implement repository and service concrete types and wire them into controllers.
3. Implement server startup in `cmd/server/main.go` using `internal/server.NewRouter()` and `config.Load()`.
4. Add tests for controllers, services, and repository implementations.

This README is intentionally short — expand with examples and contributor notes as you start implementing.
