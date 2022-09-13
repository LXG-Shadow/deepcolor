package transform

import (
	"encoding/json"
	"fmt"
	"github.com/aynakeya/deepcolor/util"
	"gotest.tools/assert"
	"regexp"
	"testing"
)

var testString = "【喵萌奶茶屋】★07月新番★[莉可丽丝/Lycoris Recoil][05][1080p][简繁内封][招募翻译校对][MKV]"

func TestRegExpTranslator(t *testing.T) {
	trans := NewRegExpReplacer(
		regexp.MustCompile(`【喵萌奶茶屋】★07月新番★\[莉可丽丝/Lycoris Recoil]\[(.*)]\[1080p]\[简繁内封]\[招募翻译校对]\[MKV]`),
		"$1",
	)
	assert.Equal(t, "05", trans.MustApply(testString))
}

func TestRegExpTranslator_Marshalling(t *testing.T) {
	trans := NewRegExpReplacer(
		regexp.MustCompile(`【喵萌奶茶屋】★07月新番★\[莉可丽丝/Lycoris Recoil]\[(.*)]\[1080p]\[简繁内封]\[招募翻译校对].*`),
		"$1",
	)
	data, err := util.MarshalIndentUnescape(trans, "", "  ")
	if err != nil {
		t.Fatalf("Marshlling failed")
	}
	var trans1 RegExpReplacer
	fmt.Println(data)
	err = json.Unmarshal([]byte(data), &trans1)
	if err != nil {
		t.Fatalf("Unmarshlling failed")
		return
	}
	assert.Equal(t, "05", trans1.MustApply(testString))
	trans2, err := UnmarshalTranslator([]byte(data))
	if err != nil {
		t.Fatalf("fail to unmarshal using UnmarshalTranslator, %s", err)
	}
	assert.Equal(t, "05", trans2.MustApply(testString))
}
