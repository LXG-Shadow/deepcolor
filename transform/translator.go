package transform

import (
	"encoding/json"
	"reflect"
)

type Translator interface {
	GetType() string
	Apply(value interface{}) (interface{}, error)
	MustApply(value interface{}) interface{}
}

type BaseTranslator struct {
	Type string
}

func (t BaseTranslator) GetType() string {
	return t.Type
}

var _RegisteredTranslator = map[string]Translator{}

func RegisterTranslator(name string, translator Translator) {
	_RegisteredTranslator[name] = translator
	//for _, translator := range translators {
	//
	//	rType := reflect.TypeOf(translator)
	//	if rType.Kind() == reflect.Ptr {
	//		_RegisteredTranslator[translator.GetType()] = reflect.New(rType.Elem()).Interface().(Translator)
	//	} else {
	//		_RegisteredTranslator[translator.GetType()] = translator
	//	}
	//}
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
	v := reflect.New(reflect.TypeOf(t).Elem()).Interface().(Translator)
	err := json.Unmarshal(data, v)
	return v, err
}

func applyTranslator(src, dest Value, translator Translator) error {
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
