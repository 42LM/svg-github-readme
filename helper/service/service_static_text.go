// Package service implements the actual service functions that deliver svg files.
package service

import (
	"context"
	"strconv"
	"time"
)

type StaticTextInput struct {
	Color      string
	FontFamily string
	FontSize   string
	Height     uint16
	Text       string
}

// StaticText returns the svg template static_text.
// Templates can be found in `../svgtemplate/svg_templates/...`
func (s *service) StaticText(ctx context.Context) error {
	fontSizeStr := s.queryParams["font_size"]
	var height int
	var err error
	if fontSizeStr != "" {
		height, err = strconv.Atoi(fontSizeStr)
		if err != nil {
			return err
		}
	} else {
		fontSizeStr = "25"
		height, err = strconv.Atoi(fontSizeStr)
		if err != nil {
			return err
		}
	}

	color := s.queryParams["color"]
	if color == "" {
		color = "000000"
	}
	fontFamily := s.queryParams["font_family"]
	if fontFamily == "" {
		fontFamily = "Open Sans"
	}
	text := s.queryParams["text"]
	if text == "" {
		text = "hello%20world"
	}

	data := StaticTextInput{
		Color:      color,
		FontFamily: fontFamily,
		FontSize:   fontSizeStr,
		Height:     uint16(height + 30),
		Text:       text,
	}

	return s.templates.ExecuteTemplate(s.responseWriter, "static_text.gosvg", data)
}

func (mw loggingMiddleware) StaticText(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		mw.logger.Info(
			"service invocation",
			"method", "StaticText",
			"took", (time.Since(begin) / 1e6).String(),
		)
	}(time.Now())
	return mw.next.StaticText(ctx)
}
