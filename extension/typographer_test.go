package extension

import (
	"testing"

	"github.com/abdullahiqbal1996/goldmark"
	"github.com/abdullahiqbal1996/goldmark/renderer/html"
	"github.com/abdullahiqbal1996/goldmark/testutil"
)

func TestTypographer(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			Typographer,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/typographer.txt", t, testutil.ParseCliCaseArg()...)
}
