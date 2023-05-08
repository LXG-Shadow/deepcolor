package parsers

import (
	"errors"
	"github.com/aynakeya/deepcolor/dphttp"
	"github.com/tidwall/gjson"
)

func GJSONParser(resp *dphttp.Response) (*gjson.Result, error) {
	body := resp.Body()
	if len(body) == 0 {
		return nil, errors.New("GJSON Parser: fail to parse empty body")
	}
	result := gjson.Parse(string(body))
	return &result, nil
}
