package main

import (
	"fmt"
	"github.com/aynakeya/deepcolor"
	"github.com/aynakeya/deepcolor/transform"
	"log"
	"regexp"
)

var cookies = map[string]string{
	"cookie": "XSRF-TOKEN=eyJpdiI6IndnaFVBSEIzNWFseFpNc0diK0RvUVE9PSIsInZhbHVlIjoicUxJcDVZaDNTV0xPeEc4cmJ4YmZLOTI0SFFcL1dFZlI2b2cwY25kTXpYaE5rNnFvenJGa0dBTTArRE9OUldDTlIiLCJtYWMiOiI1Y2M5ZWU3Mjk3YzMxMGU1OTY2MjgzNWFkNWNhMTBiNzI1YTNiYmIzNjc0YjEwNTAwMTAyYzgzMzkwYjRhMTRjIn0%3D; glidedsky_session=eyJpdiI6Ik1DRnBNV0FGTEFYcnNUOWtiMTJvM3c9PSIsInZhbHVlIjoiWnR5UVU0RVVzVEl5Mkl2RVBMUEJhQUQrZFZvb0xTSVVCT3dHcm5vWEkwZ21zcWdPRWNTWVFPcEtpSmg5Y1k5QyIsIm1hYyI6IjJiY2NhZDA5ZjRmNmQxNTQzZWM5NWU0MjQzYjAyODVjNDNkZDRmMjQxN2UyYTk4MzljNzdkNTU5OGYzY2ExOTIifQ%3D%3D;",
}

func main() {
	resp := deepcolor.Get(&deepcolor.Request{
		Url:    "http://www.glidedsky.com/level/web/crawler-basic-1",
		Header: cookies,
	})
	if resp == nil {
		log.Fatal("fail to get response")
	}
	t := deepcolor.Tentacle{
		Parser: deepcolor.NewHTMLParser(),
		ValueMapper: map[string]deepcolor.TentacleMapper{
			"Value": {
				Selector: deepcolor.TextSliceSelector("#app > main > div.container > div > div > div > div.col-md-1"),
				Translator: transform.NewForeach(
					transform.NewPipeline(
						transform.NewRegExpFindFirst(regexp.MustCompile(`\d+`), 0),
						transform.NewCast("int"),
					),
				),
			},
		},
	}
	t.Initialize(resp)
	su := 0
	for _, v := range t.GetItems()["Value"].([]interface{}) {
		su += v.(int)
	}
	fmt.Println(su)
	fmt.Printf("%#v", t.GetItems()["Value"].([]interface{})[0])
}
