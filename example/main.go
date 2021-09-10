package main

import (
	"deepcolor"
	"fmt"
)

func main()  {
	engine := deepcolor.NewEngine()
	//engine.OnRequest(func(tentacle deepcolor.Tentacle) bool {
	//	fmt.Println(tentacle.Url)
	//	return true
	//})
	engine.OnResponse(func(result deepcolor.TentacleResult) bool {
		fmt.Println(result.GetRequest().Url,result.GetSingle(deepcolor.Item{
			Type:  deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".detail_imform_name",
					Target: deepcolor.TextTarget(),
				},
			},
		}))
		return true
	})
	for i:=20200300; i < 20200345 ; i++ {
		engine.FetchAsync(fmt.Sprintf("https://www.agefans.cc/detail/%d",i))
	}
	engine.WaitUntilFinish()
}
