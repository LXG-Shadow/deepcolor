package transform

type Foreach struct {
	BaseTranslator
	InternTrans Translator
}

func NewForeach(translator Translator) Translator {
	t := &Foreach{
		BaseTranslator: BaseTranslator{
			Type: "Foreach",
		},
		InternTrans: translator,
	}
	return t
}

func (f *Foreach) Apply(value interface{}) (interface{}, error) {
	v, ok := value.([]interface{})
	if !ok {
		return value, errorWrongSrcType("[]interface")
	}
	for index, _ := range v {
		v[index] = f.InternTrans.MustApply(v[index])
	}
	return v, nil
}

func (f *Foreach) MustApply(value interface{}) interface{} {
	v, _ := f.Apply(value)
	return v
}
