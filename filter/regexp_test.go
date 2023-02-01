package filter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestRegExpFilter(t *testing.T) {
	title := "【喵萌奶茶屋】★07月新番★[莉可丽丝/Lycoris Recoil][05][1080p][简繁内封][招募翻译校对]"
	var filter1 = NewRegExpFilter(
		regexp.MustCompile("Lycoris Recoil"),
		true,
	)
	if filter1.Check(title) != true {
		t.Errorf("Include TargetTitle failed")
	}
	filter2 := NewRegExpFilter(
		regexp.MustCompile("Lycoris Recoil"),
		false,
	)
	if filter2.Check(title) != false {
		t.Errorf("Exclude TargetTitle failed")
	}
	filter3 := NewRegExpFilter(
		regexp.MustCompile(`【喵萌奶茶屋】★07月新番★\[莉可丽丝/Lycoris Recoil]\[\d+]\[1080p]\[简繁内封]\[招募翻译校对]`),
		true,
	)
	if filter3.Check(title) != true {
		t.Errorf("Include TargetTitle all failed")
	}
}

func TestRegExpFilter_Marshalling(t *testing.T) {
	var filter1 Filter = NewRegExpFilter(regexp.MustCompile("Lycoris Recoil [0-9]*"), true)
	data, err := json.MarshalIndent(filter1, "", "  ")
	if err != nil {
		t.Fatalf("Marshlling failed")
	}
	var regF RegExpFilter
	fmt.Println(data)
	err = json.Unmarshal(data, &regF)
	if err != nil {
		t.Fatalf("Unmarshlling failed, %s", err)
		return
	}
	if regF.Expression.String() != filter1.(*RegExpFilter).Expression.String() {
		t.Fatalf("Unmarshlling field not match")
		return
	}
	if regF.Expression.FindString("b9834hgbsaLycoris Recoil 33H%$43") != "Lycoris Recoil 33" {
		t.Fatalf("Unmarshlling field not match")
		return
	}
	f1, err := UnmarshalFilter([]byte(data))
	if err != nil {
		t.Fatalf("Unmarshlling using UnmarshalFilter failed %s", err)
	}
	fmt.Println(f1)
	if reflect.TypeOf(f1).String() != reflect.TypeOf(filter1).String() {
		t.Fatalf("UnmarshalFilter fail to construct correct filter real %s should be %s", reflect.TypeOf(f1).String(), reflect.TypeOf(filter1).String())
	}
}
