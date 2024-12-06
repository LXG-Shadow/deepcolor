package formatter

import "github.com/tidwall/gjson"

type jsonValueGetter struct {
	result gjson.Result
}

func (j *jsonValueGetter) Get(expression string) interface{} {
	return j.result.Get(expression).Value()
}

func NewJsonValueGetter(data string) IValueGetter {
	return &jsonValueGetter{
		result: gjson.Parse(data),
	}
}
