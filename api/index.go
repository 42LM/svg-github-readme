package handler

import (
	"net/http"

	"svg-github-readme/svgtemplate"
)

type QueryParams struct {
	Text string
}

func GenerateSVG(w http.ResponseWriter, r *http.Request) {
	qp := QueryParams{
		r.URL.Query().Get(string("text")),
	}

	svgTpls, err := svgtemplate.GetSVGTemplates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	svgTpls.ExecuteTemplate(w, "animated_text.gosvg", qp)
}
