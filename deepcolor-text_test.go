package deepcolor

import (
	"github.com/aynakeya/deepcolor/transform"
	"gotest.tools/assert"
	"log"
	"regexp"
	"testing"
)

type TextInfoStruct struct {
	X string
}

func TestFetchText(t *testing.T) {
	tenc := Tentacle{
		Parser: &ParserRegexp{},
		ValueMapper: map[string]*Selector{
			"X": RegExpSelector("<title>.*</title>"),
		},
		Transformers: []*transform.Transformer{
			{
				"X",
				"X",
				transform.NewRegExpReplacer(regexp.MustCompile("</?title>"), ""),
			},
		},
	}
	err := tenc.Initialize(QuickGet("https://crawler-test.com/", nil))
	if err != nil {
		log.Fatal(err)
		return
	}
	var v TextInfoStruct
	tenc.ExtractAndTransform(&v)
	assert.Equal(t, "Crawler Test Site", v.X)
}
