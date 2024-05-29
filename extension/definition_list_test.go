package extension

import (
	"testing"

	"github.com/abdullahiqbal1996/goldmark"
	"github.com/abdullahiqbal1996/goldmark/renderer/html"
	"github.com/abdullahiqbal1996/goldmark/testutil"
)

func TestDefinitionList(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			DefinitionList,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/definition_list.txt", t, testutil.ParseCliCaseArg()...)
}
