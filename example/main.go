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
	//engine.SetBurst(1)
	//engine.SetPeriod(time.Second * 1)
	engine.SetMaxConnection(5)
	engine.OnResponse(func(result *deepcolor.TentacleResult) bool {
		//a := rand.Intn(10)
		//fmt.Printf("Sleep %d, Get %s\n", a, result.Request.Url)
		//time.Sleep(time.Second * time.Duration(a))
		count++
		fmt.Println(result.GetRequest().Url, result.GetSingle(deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.TextSelector(".detail_imform_name"),
				},
			},
		}))
		return true
	})
	for i := 20200300; i < 20200345; i++ {
		engine.FetchAsync(fmt.Sprintf("https://www.agemys.com/detail/%d", i))
	}
	engine.WaitUntilFinish()
	fmt.Println(count)
}
