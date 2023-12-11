// Package service implements the actual service functions that deliver svg files.
package service

import (
	"context"
	"strconv"
	"time"
)

type AnimatedTextInput struct {
	AnimationHeight string
	Color           string
	FontFamily      string
	FontSize        string
	Height          uint16
	Text            string
}

// AnimatedText returns the svg template animated_text.
// Templates can be found in `../svgtemplate/svg_templates/...`
func (s *service) AnimatedText(ctx context.Context) error {
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

	data := AnimatedTextInput{
		AnimationHeight: fontSizeStr,
		Color:           color,
		FontFamily:      fontFamily,
		FontSize:        fontSizeStr,
		Height:          uint16(height + 30),
		Text:            text,
	}

	return s.templates.ExecuteTemplate(s.responseWriter, "animated_text.gosvg", data)
}

func (mw loggingMiddleware) AnimatedText(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		mw.logger.Info(
			"service invocation",
			"method", "AnimatedText",
			"took", (time.Since(begin) / 1e6).String(),
		)
	}(time.Now())
	return mw.next.AnimatedText(ctx)
}
