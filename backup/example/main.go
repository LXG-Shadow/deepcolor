package main

import (
	"fmt"
	"github.com/aynakeya/deepcolor"
	"github.com/aynakeya/deepcolor/backup"
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
	engine := backup.NewEngine(5)
	//engine.SetBurst(1)
	//engine.SetPeriod(time.Second * 1)
	//engine.SetMaxConnection(5)
	dp := backup.Deepcolor{
		ReqFunc:     deepcolor.Get,
		Requester:   &agerequester{id: 20200300},
		ReqHandler:  nil,
		RespHandler: nil,
		Tentacles: []*backup.Tentacle{
			{
				Parser: backup.NewHTMLParser(),
				ValueMapper: map[string]*backup.TentacleMapper{
					"Title": backup.SelectorText(".detail_imform_name").ToMapper(),
				},
				Handlers: []backup.TentacleHandler{
					func(tentacle *backup.Tentacle) {
						fmt.Println(tentacle.GetItems()["Title"])
					},
				},
			},
		},
	}
	engine.StartParallel(&dp)
	fmt.Println("1233")
	engine.WaitUntilFinish()
	fmt.Println("finished")
}
