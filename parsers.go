package deepcolor

import (
	"github.com/aynakeya/deepcolor/dphttp"
	"github.com/aynakeya/deepcolor/dphttp/parsers"
	"github.com/tidwall/gjson"
)

func ParserGJson(resp *dphttp.Response) (*gjson.Result, error) {
	return parsers.GJSONParser(resp)
}

func ParserText(resp *dphttp.Response) (string, error) {
	return parsers.TextParser(resp)
}
