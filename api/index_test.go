package handler_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	handler "svg-github-readme/api"
)

func Test_GenerateSVG(t *testing.T) {
	text := "luke skywalker"

	query := fmt.Sprintf(
		"/api?text=%s",
		url.QueryEscape(text),
	)

	req := httptest.NewRequest(http.MethodGet, query, nil)
	w := httptest.NewRecorder()

	handler.GenerateSVG(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error("unexpected error")
	}

	if string(data) != fixture {
		t.Errorf("\n\nexpected: %v\nactual: %v\n", fixture, string(data))
	}
}

const fixture = "<svg width=\"100%\" height=\"100%\" viewBox=\"30 -50 600 500\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" version=\"1.1\">\n <path id=\"path\">\n\t\t<animate attributeName=\"d\" from=\"m0,110 h0\" to=\"m0,110 h1100\" dur=\"6.8s\" begin=\"0s\" repeatCount=\"indefinite\"/>\n\t</path>\n\t<text font-size=\"26\" font-family=\"Montserrat\" fill='#3081D0'>\n\t\t<textPath xlink:href=\"#path\">luke skywalker</textPath>\n\t</text>\n</svg>\n"
