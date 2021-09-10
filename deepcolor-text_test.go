package deepcolor

import (
	"fmt"
	"testing"
)



func TestFetchText(t *testing.T) {
	ten := TentacleHTML("https://crawler-test.com/", "utf-8")
	result, err := Fetch(ten, func(uri string, header map[string]string) string {
		return Get(uri, header).String()
	},nil,nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	singleTestItem2 := Item{
		Type: ItemTypeSingle,
		Rules: []ItemRule{
			{
				Selector:     "<title>.*</title>",
				Target:       RegExpTarget(),
				Substitution: map[string]string{
					"</?title>":"",
				},
			},
		},
	}
	if result.GetSingle(singleTestItem2) != "Crawler Test Site"{
		t.Errorf("Fail, should be %s","Crawler Test Site")
	}
}