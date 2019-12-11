package main

import (
	"context"
	"fmt"
	"reflect"
)

// Reply ..
type Reply struct {
	r string
}

// Args ..
type Args struct {
	A int
	B string
}

// Handlers ..
type Handlers struct{}

// Foo ..
func (h Handlers) Foo(ctx context.Context, arg *Args, reply *Reply) error {
	fmt.Println("Foo is calling with", arg.A, arg.B)

	reply.r = "Hello Norman, this is the Foo result"

	return nil
}

// Goo ..
func (h *Handlers) Goo(ctx context.Context, arg *Args, reply *Reply) error {
	fmt.Println("Goo is calling with", arg.A, arg.B)

	reply.r = "Hello Norman, this is the Goo result"

	return nil
}

func main() {
	var h Handlers

	ctx := context.Background()
	arg := Args{0, "Norman"}
	var reply Reply

	fnFoo := reflect.TypeOf(h).Method(0)
	fnFoo.Func.Call([]reflect.Value{
		reflect.ValueOf(h),
		reflect.ValueOf(ctx),
		reflect.ValueOf(&arg),
		reflect.ValueOf(&reply),
	})

	fmt.Println(reply.r)

	fmt.Println()
	a := "aaa"
	if reflect.TypeOf(&a).Kind() == reflect.Ptr {
		fmt.Println("&a is a pointer")
	} else {
		fmt.Println("&a is not a pointer")
	}

	fmt.Println()
	if reflect.TypeOf(&h).Method(0).Type.Kind() == reflect.Ptr {
		fmt.Println("&h is a pointer")
	} else {
		fmt.Println("&h is not a pointer")
	}

	fmt.Println("reflect.TypeOf(h).NumMethod()", reflect.TypeOf(h).NumMethod())
	fmt.Println("reflect.TypeOf(&h).NumMethod()", reflect.TypeOf(&h).NumMethod())
}
