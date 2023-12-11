// Package handler implements a single HTTP handler
// that is being used by the Go Runtime of Vercel
package handler

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"svg-github-readme/internal/query"
	"svg-github-readme/internal/service"
	"svg-github-readme/internal/svgtemplate"
)

// GenerateSVG is the single handler being triggered by vercel.
// It generates an svg according to the given `type`.
func GenerateSVG(w http.ResponseWriter, r *http.Request) {
	// create service
	qp := query.GetQueryParams(r)
	svgTemplates, err := svgtemplate.GetSVGTemplates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	svc := service.New(&service.ServiceConfig{
		Logger:         slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		QueryParams:    qp,
		ResponseWriter: w,
		Templates:      svgTemplates,
	})

	// type switch for svg type
	// no type given leads to creation of a default error svg.
	switch qp["type"] {
	case "animated_text":
		svc.AnimatedText(context.Background())
	default:
		svc.Error(context.Background())
	}
}
