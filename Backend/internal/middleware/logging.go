package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// Basic middleware chain placeholder. Use chi/middleware or your own implementations.
func DefaultMiddlewares() func(next http.Handler) http.Handler {
	// For now return chi's default RequestLogger or a noop. Replace as needed.
	return middleware.RequestID
}

// Example custom logging can be added here when needed.
