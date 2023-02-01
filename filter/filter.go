package filter

import (
	"encoding/json"
	"errors"
	"github.com/aynakeya/deepcolor/transform"
	"reflect"
)

type Filter interface {
	GetType() string
	Check(value interface{}) bool // return true if the meta should be keep
}

type BaseFilter struct {
	Type string
}

func (f *BaseFilter) GetType() string {
	return f.Type
}

var _RegisteredFilter = map[string]Filter{}

func RegisterFilter(filters ...Filter) {
	for _, filter := range filters {
		_RegisteredFilter[filter.GetType()] = filter
	}
}

func UnmarshalFilter(data []byte) (Filter, error) {
	var f BaseFilter
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	t, ok := _RegisteredFilter[f.Type]
	if !ok {
		return nil, errors.New("filter not found")
	}
	v := reflect.New(reflect.TypeOf(t).Elem()).Interface().(Filter)
	err := json.Unmarshal(data, v)
	return v, err
}

type StructFilter struct {
	Target transform.Field
	Filter Filter
}

func (f *StructFilter) Check(value interface{}) bool {
	v := f.Target.GetValue(value)
	return f.Filter.Check(v.Interface())
}

func (f *StructFilter) UnmarshalJSON(data []byte) error {
	type Tmp StructFilter
	aux := &struct {
		*Tmp
		Filter json.RawMessage
	}{
		Tmp: (*Tmp)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var err error
	f.Filter, err = UnmarshalFilter(aux.Filter)
	return err
}
