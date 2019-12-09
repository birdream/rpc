package main

import "fmt"
import "reflect"

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
	reflect.Type().Method.Func.Call([]reflect.Value{
		reflect.ValueOf(4),
	})
}
