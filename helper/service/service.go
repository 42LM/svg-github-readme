// Package service implements the actual service functions that deliver svg files
package service

import (
	"context"
	"html/template"
	"log/slog"
	"net/http"
)

type Service interface {
	AnimatedText(ctx context.Context) error
	Error(ctx context.Context) error
}

type service struct {
	queryParams    map[string]string
	responseWriter http.ResponseWriter
	templates      *template.Template
}

var _ Service = (*service)(nil)

// ServiceConfig contains the configuration params of the service.
type ServiceConfig struct {
	Logger         *slog.Logger
	QueryParams    map[string]string
	ResponseWriter http.ResponseWriter
	Templates      *template.Template
}

// New returns a service with middleware wired in.
func New(config *ServiceConfig) Service {
	var svc Service
	svc = &service{
		queryParams:    config.QueryParams,
		responseWriter: config.ResponseWriter,
		templates:      config.Templates,
	}
	svc = LoggingMiddleware(config.Logger)(svc)
	return svc
}
