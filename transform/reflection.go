package transform

import (
	"reflect"
	"strings"
)

// todo https://stackoverflow.com/questions/47187680/how-do-i-change-fields-a-slice-of-structs-using-reflect

type Field string

func (t Field) GetValue(inst interface{}) reflect.Value {
	v := reflect.ValueOf(inst)
	for _, f := range strings.Split(string(t), ".") {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		v = v.FieldByName(f)
	}
	return v
}

func (t Field) GetValueE(inst interface{}) (reflect.Value, bool) {
	v := reflect.ValueOf(inst)
	for _, f := range strings.Split(string(t), ".") {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			return v, false
		}
		v = v.FieldByName(f)
		if !v.IsValid() {
			return v, false
		}
	}
	return v, true
}

func SetFieldValue(src interface{}, dst reflect.Value) {
	if dst.Kind() == reflect.Ptr {
		dst = dst.Elem()
	}
	if dst.Kind() != reflect.Slice {
		dst.Set(reflect.ValueOf(src))
		return
	}
	asrc, ok := src.([]interface{})
	if !ok {
		return
	}
	s := reflect.New(dst.Type()).Elem()
	for i := 0; i < len(asrc); i++ {
		s = reflect.Append(s, reflect.ValueOf(asrc[i]))
	}
	dst.Set(s)
}
