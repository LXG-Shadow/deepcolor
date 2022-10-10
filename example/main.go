package main

import (
	"fmt"
	"github.com/aynakeya/deepcolor"
)

type agerequester struct {
	id int
}

func (r *agerequester) Next(response *deepcolor.Response) *deepcolor.Request {
	if r.id >= 20200445 {
		return nil
	}
	req := &deepcolor.Request{
		Url:     fmt.Sprintf("https://www.agemys.com/detail/%d", r.id),
		Charset: "utf-8",
	}
	r.id++
	return req
}

func main() {
	engine := deepcolor.NewEngine(5)
	//engine.SetBurst(1)
	//engine.SetPeriod(time.Second * 1)
	//engine.SetMaxConnection(5)
	dp := deepcolor.Deepcolor{
		ReqFunc:     deepcolor.Get,
		Requester:   &agerequester{id: 20200300},
		ReqHandler:  nil,
		RespHandler: nil,
		Tentacles: []*deepcolor.Tentacle{
			{
				Parser: deepcolor.NewHTMLParser(),
				ValueMapper: map[string]*deepcolor.TentacleMapper{
					"Title": deepcolor.TextSelector(".detail_imform_name").ToMapper(),
				},
				Handlers: []deepcolor.TentacleHandler{
					func(tentacle *deepcolor.Tentacle) {
						fmt.Println(tentacle.GetItems()["Title"])
					},
				},
			},
		},
	}
	engine.Add(&dp)
	fmt.Println("1233")
	engine.WaitUntilFinish()
	fmt.Println("finished")
}
