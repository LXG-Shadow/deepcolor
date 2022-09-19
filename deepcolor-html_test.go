package deepcolor

import (
	"fmt"
	"github.com/aynakeya/deepcolor/transform"
	"gotest.tools/assert"
	"log"
	"regexp"
	"testing"
)

type InfoStruct struct {
	X string
	Y []string
	Z string

	B []string
}

func TestFetch(t *testing.T) {
	tenc := Tentacle{
		Parser: &ParserHTML{},
		ValueMapper: map[string]TentacleMapper{
			"X":   NewTentacleSelector(TextSelector("#logo")),
			"X.A": NewTentacleSelector(TextSelector("#logo")),
			"A":   NewTentacleSelector(TextSelector("#logo")),
			"Y":   NewTentacleSelector(TextSliceSelector("body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a")),
			"Z":   NewTentacleSelector(TextSelector("body > div:nth-child(2) > div > h1")),
			"B":   NewTentacleSelector(AttributeSliceSelector("body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a", "href")),
		},
		Transformers: []*transform.Transformer{
			{
				"X",
				"X",
				transform.NewRegExpReplacer(regexp.MustCompile("two"), "three"),
			},
		},
	}
	err := tenc.Initialize(QuickGet("https://crawler-test.com/", nil))
	if err != nil {
		log.Fatal(err)
		return
	}
	var s InfoStruct
	_ = tenc.ExtractAndTransform(&s)
	assert.Equal(t, "Crawler Test three point oh!", s.X)
	assert.Equal(t, "Crawler Test Site", s.Z)
	fmt.Println(s.Y)
	fmt.Println(s.B)
}
