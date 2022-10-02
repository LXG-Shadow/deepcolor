package transform

import (
	"encoding/json"
	"reflect"
)

type Translator interface {
	BaseType() string
	Apply(value interface{}) (interface{}, error)
	MustApply(value interface{}) interface{}
}

type BaseTranslator struct {
	Type string
}

func (t BaseTranslator) BaseType() string {
	return t.Type
}

func (b *BaseTranslator) Extend(t Translator) {
	rType := reflect.TypeOf(t)
	if rType.Kind() == reflect.Ptr {
		b.Type = rType.Elem().Name()
	} else {
		b.Type = rType.Name()
	}
}

var _RegisteredTranslator = map[string]reflect.Type{}

func RegisterTranslator(translators ...Translator) {
	for _, translator := range translators {
		rType := reflect.TypeOf(translator)
		if rType.Kind() == reflect.Ptr {
			_RegisteredTranslator[rType.Elem().Name()] = rType.Elem()
		} else {
			_RegisteredTranslator[rType.Name()] = rType
		}
	}
}

func UnmarshalTranslator(data []byte) (Translator, error) {
	var f BaseTranslator
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	t, ok := _RegisteredTranslator[f.Type]
	if !ok {
		return nil, errorTranslatorNotFound
	}
	v := reflect.New(t).Interface().(Translator)
	err := json.Unmarshal(data, v)
	return v, err
}

func applyTranslator(src, dest reflect.Value, translator Translator) error {
	value, err := translator.Apply(src.Interface())
	if err != nil {
		return err
	}
	dest.Set(reflect.ValueOf(value))
	return nil
}

func Transform(value interface{}, translator Translator, src Field, dest Field) error {
	return applyTranslator(src.GetValue(value), dest.GetValue(value), translator)
}
