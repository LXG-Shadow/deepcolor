package main

import (
	"fmt"
	"reflect"
)

type BaseTranslator struct {
	Type string
}

type X struct {
	BaseTranslator
	Y string
}

func main() {
	x := &X{
		BaseTranslator: BaseTranslator{Type: "123"},
		Y:              "asdf",
	}
	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(reflect.TypeOf(x))
}
