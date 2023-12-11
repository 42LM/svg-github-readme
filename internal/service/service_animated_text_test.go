package service_test

import (
	"context"
	"io"
	"log/slog"
	"net/http/httptest"
	"testing"

	"svg-github-readme/internal/service"
	"svg-github-readme/internal/svgtemplate"
)

func Test_Service_AnimatedText(t *testing.T) {
	testCases := map[string]struct {
		queryParams map[string]string

		expBody string
		expErr  error
	}{
		"all valid params": {
			queryParams: map[string]string{
				"color":       "EEEEEE",
				"font_family": "Arial",
				"font_size":   "20",
				"text":        "hello world",
			},
			expBody: animated_text_svg,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// take a look at the general URL form:
			// https://pkg.go.dev/net/url#:~:text=%5Bscheme%3A%5D%5B//%5Buserinfo%40%5Dhost%5D%5B/%5Dpath%5B%3Fquery%5D%5B%23fragment%5D
			svgTemplates, _ := svgtemplate.GetSVGTemplates()

			rec := httptest.NewRecorder()
			s := service.New(&service.ServiceConfig{
				Logger:         slog.Default(),
				QueryParams:    tc.queryParams,
				ResponseWriter: rec,
				Templates:      svgTemplates,
			})

			err := s.AnimatedText(context.Background())

			// Get the captured response from the ResponseRecorder
			resp := rec.Result()

			// Read the response body
			body, _ := io.ReadAll(resp.Body)

			if tc.expBody != string(body) {
				t.Errorf("\nexpected: %v\nactual: %v\n", tc.expBody, string(body))
			}

			if tc.expErr != err {
				t.Errorf("\nexpected: %v\nactual: %v\n", tc.expErr, err)
			}
		})
	}
}

const (
	animated_text_svg = "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" version=\"1.1\" width=\"1280\" height=\"50\">\n <title>animated_text</title>\n <path id=\"path\" x=\"0\" y=\"0\">\n   <animate attributeName=\"d\" from=\"m0,20 h0\" to=\"m0,20 h1100\" dur=\"6.8s\" begin=\"0s\" repeatCount=\"indefinite\"/>\n </path>\n  <text x=\"0\" y=\"0\" font-size=\"20\" font-family=\"Arial\" fill=\"#EEEEEE\">\n    <textPath xlink:href=\"#path\">hello world</textPath>\n  </text>\n</svg>\n"
)
