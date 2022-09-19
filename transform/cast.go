package transform

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
)

type Cast struct {
	BaseTranslator
	ToType string
}

func NewCast(destType string) Translator {
	t := &Cast{
		ToType: destType,
	}
	t.Extend(t)
	return t
}

func (c *Cast) Apply(value interface{}) (interface{}, error) {
	switch c.ToType {
	case "string":
		return cast.ToStringE(value)
	case "bool":
		return cast.ToBoolE(value)
	case "int":
		return cast.ToIntE(value)
	}
	return value, errors.New(fmt.Sprintf("%s not support in casting", c.ToType))
}

func (c *Cast) MustApply(value interface{}) interface{} {
	v, _ := c.Apply(value)
	return v
}
