package deepcolor

import (
	"regexp"
)

type ParserRegexp struct {
	data string
}

func NewRegexpParser() ResponseParser {
	return &ParserRegexp{}
}

func (p *ParserRegexp) Initialize(resp *Response) error {
	body := resp.Body()
	if len(body) == 0 {
		return ErrorEmptyBody
	}
	p.data = string(body)
	return nil
}

func (p *ParserRegexp) Get(selector *Selector) interface{} {
	if selector.Array {
		return p.GetValues(selector)
	}
	return p.GetValue(selector)
}

func (p *ParserRegexp) GetValue(selector *Selector) interface{} {
	if selector.Path == "" {
		return ""
	}
	return regexp.MustCompile(selector.Path).FindString(p.data)
}

func (p *ParserRegexp) GetValues(selector *Selector) []interface{} {
	v0 := make([]interface{}, 0)
	if selector.Path == "" {
		return v0
	}
	v1 := regexp.MustCompile(selector.Path).FindAllString(p.data, -1)
	for _, v := range v1 {
		v0 = append(v0, v)
	}
	return v0
}
