package service_test

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http/httptest"
	"testing"

	"svg-github-readme/helper/service"
	"svg-github-readme/helper/svgtemplate"
)

func Test_Service_StaticText(t *testing.T) {
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
			expBody: static_text_svg,
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

			err := s.StaticText(context.Background())

			// Get the captured response from the ResponseRecorder
			resp := rec.Result()

			// Read the response body
			body, _ := io.ReadAll(resp.Body)

			fmt.Printf("HM: %#v", string(body))
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
	static_text_svg = "<svg xmlns=\"http://www.w3.org/2000/svg\" version=\"1.1\" viewbox=\"0 0 100 100\" width=\"1280\" height=\"50\">\n  <title>static_text</title>\n  <text fill=\"#EEEEEE\">\n    <tspan font-size=\"20\" x=\"0\" y=\"50%\" font-family=\"Arial\" dominant-baseline=\"middle\">hello world</tspan>\n  </text>\n</svg>\n"
)
