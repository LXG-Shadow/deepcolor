package deepcolor

import (
	"fmt"
	"github.com/aynakeya/deepcolor/transform"
	"gotest.tools/assert"
	"log"
	"regexp"
	"testing"
)

type InfoStructJson struct {
	X string
	Y []string
}

func TestFetchJson(t *testing.T) {
	tenc := Tentacle{
		Parser: &ParserJson{},
		ValueMapper: map[string]*Selector{
			"X": JsonSelector("fruit"),
			"Y": JsonSliceSelector("quiz.sport.q1.options"),
		},
		Transformers: []*transform.Transformer{
			{
				"X",
				"X",
				transform.NewRegExpReplacer(regexp.MustCompile("p"), "b"),
			},
		},
	}
	err := tenc.Initialize(QuickGet("https://support.oneskyapp.com/hc/en-us/article_attachments/202761627/example_1.json", nil))
	if err != nil {
		log.Fatal(err)
		return
	}
	var s InfoStructJson
	tenc.ExtractAndTransform(&s)
	assert.Equal(t, "Abble", s.X)
	err = tenc.Initialize(QuickGet("https://support.oneskyapp.com/hc/en-us/article_attachments/202761727/example_2.json", nil))
	if err != nil {
		log.Fatal(err)
		return
	}
	tenc.ExtractAndTransform(&s)
	fmt.Println(s.Y)
}
