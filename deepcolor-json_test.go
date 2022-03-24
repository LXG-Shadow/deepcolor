package deepcolor

import (
	"fmt"
	"testing"
)

func TestFetchJson(t *testing.T) {
	ten := TentacleJson("https://support.oneskyapp.com/hc/en-us/article_attachments/202761627/example_1.json", "utf-8")
	result, err := Fetch(ten, Get, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	singleTestItem2 := Item{
		Type: ItemTypeSingle,
		Rules: []ItemRule{
			{
				Selector: JsonSelector("fruit"),
				Substitution: map[string]string{
					"p": "b",
				},
			},
		},
	}
	fmt.Println(result.GetSingle(singleTestItem2))
	if result.GetSingle(singleTestItem2) != "Abble" {
		t.Errorf("Fail, should be %s", "Crawler Test Site")
	}

	ten = TentacleJson("https://support.oneskyapp.com/hc/en-us/article_attachments/202761727/example_2.json", "utf-8")
	result, err = Fetch(ten, Get, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	listTestItem := Item{
		Type: ItemTypeList,
		Rules: []ItemRule{
			{
				Selector: JsonSelector("quiz.sport.q1.options"),
			},
		},
	}
	fmt.Println(result.GetList(listTestItem))
}
