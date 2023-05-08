package translators

import "github.com/aynakeya/deepcolor/transform"

type Setter struct {
	transform.BaseTranslator
	Value interface{}
}

func NewSetter(value interface{}) transform.Translator {
	t := &Setter{
		BaseTranslator: transform.BaseTranslator{
			Type: "Setter",
		},
		Value: value,
	}
	return t
}

func (f *Setter) Apply(value interface{}) (interface{}, error) {
	return f.Value, nil
}

func (f *Setter) MustApply(value interface{}) interface{} {
	v, _ := f.Apply(value)
	return v
}
