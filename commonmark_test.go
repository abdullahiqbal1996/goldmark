package goldmark_test

import (
	"encoding/json"
	"os"
	"testing"

	. "github.com/abdullahiqbal1996/goldmark"
	"github.com/abdullahiqbal1996/goldmark/renderer/html"
	"github.com/abdullahiqbal1996/goldmark/testutil"
)

type commonmarkSpecTestCase struct {
	Markdown  string `json:"markdown"`
	HTML      string `json:"html"`
	Example   int    `json:"example"`
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
	Section   string `json:"section"`
}

func TestSpec(t *testing.T) {
	bs, err := os.ReadFile("_test/spec.json")
	if err != nil {
		panic(err)
	}
	var testCases []commonmarkSpecTestCase
	if err := json.Unmarshal(bs, &testCases); err != nil {
		panic(err)
	}
	cases := []testutil.MarkdownTestCase{}
	nos := testutil.ParseCliCaseArg()
	for _, c := range testCases {
		shouldAdd := len(nos) == 0
		if !shouldAdd {
			for _, no := range nos {
				if c.Example == no {
					shouldAdd = true
					break
				}
			}
		}

		if shouldAdd {
			cases = append(cases, testutil.MarkdownTestCase{
				No:       c.Example,
				Markdown: c.Markdown,
				Expected: c.HTML,
			})
		}
	}
	markdown := New(WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	))
	testutil.DoTestCases(markdown, cases, t)
}
