// Package service implements the actual service functions that deliver svg files.
package service

import (
	"log/slog"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// LoggingMiddleware takes a logger as a dependency and returns a service Middleware.
func LoggingMiddleware(logger *slog.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger *slog.Logger
	next   Service
}
