package main

import (
	"context"
	"fmt"
	"reflect"
)

func main() {
	var h Handlers1

	var (
		ok           bool
		srvInterface interface{}
		srv          *service
		mtype        *methodType
	)

	Register(Handlers1{})

	fmt.Println("----OK Register")

	srvName := "Handlers1"
	mName := "Foo"

	srvInterface, ok = serviceMap.Load(srvName)
	if !ok {
		fmt.Println("It is not ok for loading the srv interface", srvName)
	}

	srv, ok = srvInterface.(*service)
	if !ok {
		fmt.Println("It is not ok for loading the srv", srvName)
	}

	mtype, ok = srv.methods[mName]
	if !ok {
		fmt.Println("It is not ok for loading the srv method", mName)
	}

	// ctx := context.Background()

	ctx := context.Background()
	args := &Args{0, "Hey"}
	reply := &Reply{}

	mtype.Method.Func.Call([]reflect.Value{
		reflect.ValueOf(h),
		reflect.ValueOf(ctx),
		reflect.ValueOf(args),
		reflect.ValueOf(reply),
	})

	fmt.Println(reply.r)
}
