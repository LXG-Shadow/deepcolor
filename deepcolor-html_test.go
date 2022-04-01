package deepcolor

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestFetch(t *testing.T) {
	ten := TentacleHTML("https://crawler-test.com/", "utf-8")
	result, err := Fetch(ten, Get, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	singleTestItem := Item{
		Type: ItemTypeSingle,
		Rules: []ItemRule{
			{
				Selector: TextSelector("#logo"),
			},
		},
	}
	listTestItem := Item{
		Type: ItemTypeList,
		Rules: []ItemRule{
			{
				Selector: TextSelector("body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a"),
			},
		},
	}
	mapTestItem := Item{
		Type: ItemTypeMap,
		Rules: []ItemRule{
			{
				Key:      "title",
				Selector: TextSelector("body > div:nth-child(2) > div > h1"),
			},
		},
	}
	fmt.Println("\nsingle item - Text\n", result.GetSingle(singleTestItem))
	assert.Equal(t, result.GetSingle(singleTestItem), "Crawler Test two point oh!")
	fmt.Println("\nlist item - Text\n", result.GetList(listTestItem))
	//assert.Equal(t, result.GetList(listTestItem), "Crawler Test two point oh!")
	fmt.Println("\nmap item - Text\n", result.GetMap(mapTestItem))

	singleTestItem2 := Item{
		Type: ItemTypeSingle,
		Rules: []ItemRule{
			{
				Selector: TextSelector("#logo"),
			},
			{
				Selector: TextSelector("body > div:nth-child(2) > div > h1"),
			},
		},
	}

	fmt.Println("\nsingle item - Text\n", result.GetSingle(singleTestItem2))

	listTestItem2 := Item{
		Type: ItemTypeList,
		Rules: []ItemRule{
			{
				Selector: AttributeSelector(
					"body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a",
					"href"),
			},
		},
	}
	fmt.Println("\nlist item - attribute\n", result.GetList(listTestItem2))

	listTestItem3 := Item{
		Type: ItemTypeList,
		Rules: []ItemRule{
			{
				Selector: TextSelector(
					"body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a",
				),
				Substitution: map[string]string{
					"Tag": "miao",
				},
			},
			{
				Selector: AttributeSelector(
					"body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a",
					"href"),
			},
		},
	}
	fmt.Println("\nlist item - text + attribute - substitution\n", result.GetList(listTestItem3))
}
