package fuzz

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/abdullahiqbal1996/goldmark"
	"github.com/abdullahiqbal1996/goldmark/extension"
	"github.com/abdullahiqbal1996/goldmark/parser"
	"github.com/abdullahiqbal1996/goldmark/renderer/html"
	"github.com/abdullahiqbal1996/goldmark/util"
)

func fuzz(f *testing.F) {
	f.Fuzz(func(t *testing.T, orig string) {
		markdown := goldmark.New(
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
				parser.WithAttribute(),
			),
			goldmark.WithRendererOptions(
				html.WithUnsafe(),
				html.WithXHTML(),
			),
			goldmark.WithExtensions(
				extension.DefinitionList,
				extension.Footnote,
				extension.GFM,
				extension.Typographer,
				extension.Linkify,
				extension.Table,
				extension.TaskList,
			),
		)
		var b bytes.Buffer
		if err := markdown.Convert(util.StringToReadOnlyBytes(orig), &b); err != nil {
			panic(err)
		}
	})
}

func FuzzDefault(f *testing.F) {
	bs, err := os.ReadFile("../_test/spec.json")
	if err != nil {
		panic(err)
	}
	var testCases []map[string]interface{}
	if err := json.Unmarshal(bs, &testCases); err != nil {
		panic(err)
	}
	for _, c := range testCases {
		f.Add(c["markdown"])
	}
	fuzz(f)
}
