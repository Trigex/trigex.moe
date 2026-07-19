package render

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var markdown = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
)

func MarkdownHTML(source string) (string, error) {
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(source), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
