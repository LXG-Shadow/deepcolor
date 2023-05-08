package transform

import (
	"github.com/aynakeya/deepcolor/pkg/dynmarshaller"
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

var translatorRegistry = dynmarshaller.NewDynamicUnmarshaller[Translator](make(map[string]Translator))

func RegisterTranslator(translators ...Translator) {
	for _, translator := range translators {
		translatorRegistry.Register(translator.GetType(), translator)
	}
}

func UnmarshalTranslator(data []byte) (Translator, error) {
	return translatorRegistry.Unmarshal(data)
}
