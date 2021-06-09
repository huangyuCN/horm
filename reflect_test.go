package horm

import (
	"fmt"
	"reflect"
	"testing"
)

type A []int

func printInfo(t reflect.Type) {
	fmt.Printf("Kind = %s\tName = %s\n", t.Kind(), t.Name())
}

func TestReflect(t *testing.T) {
	a := &A{}
	printInfo(reflect.TypeOf(a))
	printInfo(reflect.TypeOf(a).Elem())
	printInfo(reflect.TypeOf(a).Elem().Elem())
}
