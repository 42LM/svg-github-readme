package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "svg-github-readme/api"
)

func Test_GenerateSVG(t *testing.T) {
	testCases := map[string]struct {
		query string

		expSVG string
		expErr error
	}{
		"animated text": {
			query:  "?type=animated_text&text=hello%20world&color=161A30&font_size=100&font_family=Arial",
			expSVG: animated_text_svg,
		},
		"static text": {
			query:  "?type=static_text&text=hello%20world&color=161A30&font_size=100&font_family=Arial",
			expSVG: static_text_svg,
		},
		"error": {
			query:  "?text=hello%20world",
			expSVG: error_svg,
		},
		"static text default": {
			query:  "?type=static_text&text=hello%20world",
			expSVG: default_static_text_svg,
		},
		"animated text default": {
			query:  "?type=animated_text&text=hello%20world",
			expSVG: default_animated_text_svg,
		},
	}
	for tname, tc := range testCases {
		t.Run(tname, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api"+tc.query, nil)
			w := httptest.NewRecorder()

			handler.GenerateSVG(w, req)

			res := w.Result()
			defer res.Body.Close()

			data, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error("unexpected error")
			}

			if string(data) != tc.expSVG {
				t.Errorf("\n\nexpected: %#v\nactual: %#v\n", tc.expSVG, string(data))
			}
		})
	}
}

// TODO: Better testing strategy
// maybe also use the templates instead
const (
	default_animated_text_svg = "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" version=\"1.1\" width=\"1280\" height=\"55\">\n <title>animated_text</title>\n <path id=\"path\" x=\"0\" y=\"0\">\n   <animate attributeName=\"d\" from=\"m0,25 h0\" to=\"m0,25 h1100\" dur=\"6.8s\" begin=\"0s\" repeatCount=\"indefinite\"/>\n </path>\n  <text x=\"0\" y=\"0\" font-size=\"25\" font-family=\"Open Sans\" fill=\"#000000\">\n    <textPath xlink:href=\"#path\">hello world</textPath>\n  </text>\n</svg>\n"
	animated_text_svg         = "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" version=\"1.1\" width=\"1280\" height=\"130\">\n <title>animated_text</title>\n <path id=\"path\" x=\"0\" y=\"0\">\n   <animate attributeName=\"d\" from=\"m0,100 h0\" to=\"m0,100 h1100\" dur=\"6.8s\" begin=\"0s\" repeatCount=\"indefinite\"/>\n </path>\n  <text x=\"0\" y=\"0\" font-size=\"100\" font-family=\"Arial\" fill=\"#161A30\">\n    <textPath xlink:href=\"#path\">hello world</textPath>\n  </text>\n</svg>\n"
	default_static_text_svg   = "<svg xmlns=\"http://www.w3.org/2000/svg\" version=\"1.1\" viewbox=\"0 0 100 100\" width=\"1280\" height=\"55\">\n  <title>static_text</title>\n  <text fill=\"#000000\">\n    <tspan font-size=\"25\" x=\"0\" y=\"50%\" font-family=\"Open Sans\" dominant-baseline=\"middle\">hello world</tspan>\n  </text>\n</svg>\n"
	static_text_svg           = "<svg xmlns=\"http://www.w3.org/2000/svg\" version=\"1.1\" viewbox=\"0 0 100 100\" width=\"1280\" height=\"130\">\n  <title>static_text</title>\n  <text fill=\"#161A30\">\n    <tspan font-size=\"100\" x=\"0\" y=\"50%\" font-family=\"Arial\" dominant-baseline=\"middle\">hello world</tspan>\n  </text>\n</svg>\n"
	error_svg                 = "<svg width=\"100\" height=\"100\" viewBox=\"0 0 1397 1296\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\">\n  <title>error</title>\n  <g clip-path=\"url(#clip0_25_7)\">\n    <path d=\"M698.5 166L1207.29 1047.25H189.71L698.5 166Z\" fill=\"#F31559\"/>\n    <path d=\"M645 355V802H753V355H645Z\" fill=\"#D9D9D9\"/>\n    <path d=\"M645 882V990H753V882H645Z\" fill=\"#D9D9D9\"/>\n  </g>\n  <defs>\n    <clipPath id=\"clip0_25_7\">\n      <rect width=\"1397\" height=\"1296\" fill=\"white\"/>\n    </clipPath>\n  </defs>\n</svg>\n"
)
