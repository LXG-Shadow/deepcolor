package transform

type Setter struct {
	BaseTranslator
	Value interface{}
}

func NewSetter(value interface{}) Translator {
	t := &Setter{
		BaseTranslator: BaseTranslator{
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
