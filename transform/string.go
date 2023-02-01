package transform

import (
	"fmt"
	"strings"
)

type StrCase struct {
	BaseTranslator
	Lowercase bool
}

func NewStrCase(lowercase bool) Translator {
	t := &StrCase{BaseTranslator{"StrCase"}, lowercase}
	return t
}

func (c *StrCase) Apply(value interface{}) (interface{}, error) {
	s, ok := value.(string)
	if !ok {
		return s, errorWrongSrcType("string")
	}
	if c.Lowercase {
		return strings.ToLower(s), nil
	}
	return strings.ToUpper(s), nil
}

func (c *StrCase) MustApply(value interface{}) interface{} {
	v, _ := c.Apply(value)
	return v
}

type Formatter struct {
	BaseTranslator
	Format string
}

func NewFormatter(format string) Translator {
	t := &Formatter{BaseTranslator{"Formatter"}, format}
	return t
}

func (f *Formatter) Apply(value interface{}) (interface{}, error) {
	return fmt.Sprintf(f.Format, value), nil
}

func (f *Formatter) MustApply(value interface{}) interface{} {
	v, _ := f.Apply(value)
	return v
}
