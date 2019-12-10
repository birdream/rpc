package main

import (
	"fmt"
	"reflect"
)

type T struct{}

func (t *T) Foo() {
	fmt.Println("foo")
}

func (t *T) Goo(a int) {
	fmt.Println("goo", a)
}

func main() {
	var t T

	// aa := 2
	reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})
	reflect.ValueOf(&t).MethodByName("Goo").Call([]reflect.Value{
		reflect.ValueOf(4),
	})

	// reflect.ValueOf(&t).Method.Func.Call()
	// reflect.Method.Func.Call()

	// fmt.Println(reflect.TypeOf(t).Method(1).Func)

	fmt.Println(reflect.TypeOf(&t).NumMethod())

	fmt.Println(reflect.TypeOf(&t).Method(0).Name)
	// fmt.Println(reflect.TypeOf(&t).Method(1).Name)

	// T.Method(0).Func.Call([]reflect.Value{
	// 	reflect.ValueOf(4),
	// })

	// ctx := context.Background()
	// fmt.Println(reflect.TypeOf(&t).Method(1).Name)
	// fmt.Println(reflect.TypeOf(&t).Method(1).Type)

	method1Name := reflect.TypeOf(&t).Method(1).Name

	if reflect.TypeOf(&t).Method(1).Type.Kind() == reflect.Ptr {
		fmt.Println(method1Name, " is pointer")
	} else {
		fmt.Println(method1Name, " is not a pointer")
	}

	fmt.Println()

	methodFoo := reflect.TypeOf(&t).Method(0)
	methodFoo.Func.Call([]reflect.Value{
		reflect.ValueOf(&t),
	})

	methodGoo := reflect.TypeOf(&t).Method(1)
	methodGoo.Func.Call([]reflect.Value{
		reflect.ValueOf(&t),
		reflect.ValueOf(4),
	})
}
