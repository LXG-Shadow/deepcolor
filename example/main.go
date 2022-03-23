package main

import (
	"fmt"
	"github.com/aynakeya/deepcolor"
)

func main() {
	engine := deepcolor.NewEngine()
	//engine.OnRequest(func(tentacle deepcolor.Tentacle) bool {
	//	fmt.Println(tentacle.Url)
	//	return true
	//})
	count := 0
	engine.OnResponse(func(result deepcolor.TentacleResult) bool {
		count++
		fmt.Println(result.GetRequest().Url, result.GetSingle(deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".detail_imform_name",
					Target:   deepcolor.TextTarget(),
				},
			},
		}))
		return true
	})
	for i := 20200000; i < 20200345; i++ {
		engine.FetchAsync(fmt.Sprintf("https://www.agemys.com/detail/%d", i))
	}
	engine.WaitUntilFinish()
	fmt.Println(count)
}
