package goldmark_test

import (
	"testing"

	. "github.com/abdullahiqbal1996/goldmark"
	"github.com/abdullahiqbal1996/goldmark/parser"
	"github.com/abdullahiqbal1996/goldmark/testutil"
)

func TestAttributeAndAutoHeadingID(t *testing.T) {
	markdown := New(
		WithParserOptions(
			parser.WithAttribute(),
			parser.WithAutoHeadingID(),
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/options.txt", t, testutil.ParseCliCaseArg()...)
}
