package transform

import (
	"encoding/json"
)

type Transformer struct {
	Src  Field
	Dest Field
	Step Translator
}

func NewTransformer(src Field, dest Field, step Translator) *Transformer {
	return &Transformer{
		Src:  src,
		Dest: dest,
		Step: step,
	}
}

func (r *Transformer) UnmarshalJSON(data []byte) error {
	type Tmp Transformer
	aux := &struct {
		*Tmp
		Step json.RawMessage
	}{
		Tmp: (*Tmp)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var err error
	r.Step, err = UnmarshalTranslator(aux.Step)
	return err
}

func (t *Transformer) Transform(value interface{}) error {
	return Transform(value, t.Step, t.Src, t.Dest)
}
