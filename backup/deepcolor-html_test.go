package backup

import (
	"fmt"
	"github.com/aynakeya/deepcolor/transform"
	"github.com/aynakeya/deepcolor/transform/translators"
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
		ValueMapper: map[string]*TentacleMapper{
			"X":   SelectorText("#logo").ToMapper(),
			"X.A": SelectorText("#logo").ToMapper(),
			"A":   SelectorText("#logo").ToMapper(),
			"Y":   SelectorTextSlice("body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a").ToMapper(),
			"Z":   SelectorText("body > div:nth-child(2) > div > h1").ToMapper(),
			"B":   SelectorAttributeSlice("body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a", "href").ToMapper(),
		},
		Transformers: []*transform.Transformer{
			{
				"X",
				"X",
				translators.NewRegExpReplacer(regexp.MustCompile("two"), "three"),
			},
		},
	}
	err := tenc.Initialize(quickGet("https://crawler-test.com/", nil))
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
