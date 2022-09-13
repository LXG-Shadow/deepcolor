package transform

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

type A struct {
	B string
}

type C struct {
	D *A
}

type E struct {
	F C
}

func TestField_GetValue(t *testing.T) {
	s := &E{
		F: C{
			D: &A{
				B: "final",
			},
		},
	}
	assert.Equal(t, "final", Field("F.D.B").GetValue(s).Interface())
	assert.Equal(t, "final", Field("F.D.B").GetValue(*s).Interface())
	Field("F.D.B").GetValue(s).Set(reflect.ValueOf("3333"))
	assert.Equal(t, "3333", Field("F.D.B").GetValue(s).Interface())
}
