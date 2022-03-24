package deepcolor

import (
	"fmt"
	"testing"
)

func TestFetchText(t *testing.T) {
	ten := TentacleHTML("https://crawler-test.com/", "utf-8")
	result, err := Fetch(ten, Get, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	singleTestItem2 := Item{
		Type: ItemTypeSingle,
		Rules: []ItemRule{
			{
				Selector: RegExpSelector("<title>.*</title>"),
				Substitution: map[string]string{
					"</?title>": "",
				},
			},
		},
	}
	fmt.Println(result.GetSingle(singleTestItem2) + "1")
	if result.GetSingle(singleTestItem2) != "Crawler Test Site" {
		t.Errorf("Fail, should be %s", "Crawler Test Site")
	}
}
