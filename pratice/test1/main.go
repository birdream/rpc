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

func (t T) Zoo(a *string) (res *string) {
	fmt.Println("zoo", a)

	res = a

	return
}

func newValue(t reflect.Type) interface{} {
	if t.Kind() == reflect.Ptr {
		fmt.Println("is pointer")
		return reflect.New(t.Elem()).Interface()
	}

	return reflect.New(t).Interface()
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

	method2Name := reflect.TypeOf(t).Method(0).Name
	if reflect.TypeOf(t).Method(0).Type.Kind() == reflect.Ptr {
		fmt.Println(method2Name, " is pointer")
	} else {
		fmt.Println(method2Name, " is not a pointer")
	}

	fmt.Println()

	// reflect.Typeof().Method().Func.Call Sample
	methodFoo := reflect.TypeOf(&t).Method(0)
	methodFoo.Func.Call([]reflect.Value{
		reflect.ValueOf(&t),
	})

	methodGoo := reflect.TypeOf(&t).Method(1)
	methodGoo.Func.Call([]reflect.Value{
		reflect.ValueOf(&t),
		reflect.ValueOf(4),
	})

	methodZoo := reflect.TypeOf(t).Method(0)
	str := "hello Norman"
	// var res *string

	methodZoo.Func.Call([]reflect.Value{
		reflect.ValueOf(t),
		reflect.ValueOf(&str),
	})

}
