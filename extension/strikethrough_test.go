package extension

import (
	"testing"

	"github.com/abdullahiqbal1996/goldmark"
	"github.com/abdullahiqbal1996/goldmark/renderer/html"
	"github.com/abdullahiqbal1996/goldmark/testutil"
)

func TestStrikethrough(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			Strikethrough,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/strikethrough.txt", t, testutil.ParseCliCaseArg()...)
}
