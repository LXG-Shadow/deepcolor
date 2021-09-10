package deepcolor

import (
	"fmt"
	"testing"
)



func TestFetch(t *testing.T) {
	ten := TentacleHTML("https://crawler-test.com/", "utf-8")
	result, err := Fetch(ten, func(uri string, header map[string]string) string {
		return Get(uri, header).String()
	},nil,nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	singleTestItem := Item{
		Type: ItemTypeSingle,
		Rules: []ItemRule{
			{
				Selector:     "#logo",
				Target:       TextTarget(),
			},
		},
	}
	listTestItem := Item{
		Type: ItemTypeList,
		Rules: []ItemRule{
			{
				Selector:     "body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a",
				Target:       TextTarget(),
			},
		},
	}
	mapTestItem := Item{
		Type: ItemTypeMap,
		Rules: []ItemRule{
			{
				Key: "title",
				Selector:     "body > div:nth-child(2) > div > h1",
				Target:       TextTarget(),
			},
		},
	}
	fmt.Println("\nsingle item - Text\n",result.GetSingle(singleTestItem))
	fmt.Println("\nlist item - Text\n",result.GetList(listTestItem))
	fmt.Println("\nmap item - Text\n",result.GetMap(mapTestItem))

	singleTestItem2 := Item{
		Type: ItemTypeSingle,
		Rules: []ItemRule{
			{
				Selector:     "#logo",
				Target:       TextTarget(),
			},
			{
				Selector:     "body > div:nth-child(2) > div > h1",
				Target:       TextTarget(),
			},
		},
	}

	fmt.Println("\nsingle item - Text\n",result.GetSingle(singleTestItem2))

	listTestItem2 := Item{
		Type: ItemTypeList,
		Rules: []ItemRule{
			{
				Selector:     "body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a",
				Target:       AttributeTarget("href"),
			},
		},
	}
	fmt.Println("\nlist item - attribute\n",result.GetList(listTestItem2))

	listTestItem3 := Item{
		Type: ItemTypeList,
		Rules: []ItemRule{
			{
				Selector:     "body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a",
				Target:       TextTarget(),
				Substitution: map[string]string{
					"Tag":"miao",
				},
			},
			{
				Selector:     "body > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > a",
				Target:       AttributeTarget("href"),
			},
		},
	}
	fmt.Println("\nlist item - text + attribute - substitution\n",result.GetList(listTestItem3))
}