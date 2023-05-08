package backup

import (
	"github.com/tidwall/gjson"
)

type ParserJson struct {
	result gjson.Result
}

func NewJsonParser() ResponseParser {
	return &ParserJson{}
}

func (p *ParserJson) Initialize(resp *Response) error {
	body := resp.Body()
	if len(body) == 0 {
		return ErrorEmptyBody
	}
	p.result = gjson.Parse(string(body))
	return nil
}

func (p *ParserJson) Get(selector *Selector) interface{} {
	if selector.Array {
		return p.GetValues(selector)
	}
	return p.GetValue(selector)
}

func (p *ParserJson) GetValue(selector *Selector) interface{} {
	if selector.Path == "" {
		return ""
	}
	return p.result.Get(selector.Path).String()
}

func (p *ParserJson) GetValues(selector *Selector) []interface{} {
	r0 := p.result.Get(selector.Path).Array()
	r1 := make([]interface{}, len(r0))
	for index, ele := range r0 {
		r1[index] = ele.String()
	}

	return r1
}
