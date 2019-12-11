package main

import (
	"fmt"
	"reflect"
	"sync"
)

var serviceMap sync.Map

type methodType struct {
	Method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
}

type service struct {
	name    string
	typ     reflect.Type
	rcVal   reflect.Value
	methods map[string]*methodType
}

// Register ..
func Register(handlerStruct interface{}) {
	srvName := reflect.TypeOf(handlerStruct).Name()

	// init the service
	srv := new(service)
	srv.name = srvName
	srv.rcVal = reflect.ValueOf(handlerStruct)
	srv.typ = reflect.TypeOf(handlerStruct)
	srv.methods = make(map[string]*methodType)

	// loop the method of the handler struct and map them to methods
	for i := 0; i < reflect.TypeOf(handlerStruct).NumMethod(); i++ {
		method := reflect.TypeOf(handlerStruct).Method(i)
		mname := method.Name
		mtype := method.Type
		mreply := mtype.Out(0) // reply must be an error not anything else

		srv.methods[mname] = &methodType{
			Method:    reflect.TypeOf(handlerStruct).Method(i),
			ArgType:   mtype,
			ReplyType: mreply,
		}
	}

	fmt.Println("reflect.TypeOf(handlerStruct).NumMethod()", reflect.TypeOf(handlerStruct).NumMethod())
	fmt.Println("reflect.TypeOf(&handlerStruct).NumMethod()", reflect.TypeOf(&handlerStruct).NumMethod())
	fmt.Println("reflect.PtrTo(reflect.TypeOf(handlerStruct)).NumMethod()", reflect.PtrTo(reflect.TypeOf(handlerStruct)).NumMethod())

	// fmt.Println("reflect.TypeOf(&handlerStruct).NumMethod()", reflect.TypeOf(handlerStruct).Method(0))
	// for i := 0; i < reflect.TypeOf(&handlerStruct).NumMethod(); i++ {
	// 	method := reflect.TypeOf(&handlerStruct).Method(i)
	// 	mname := method.Name
	// 	mtype := method.Type
	// 	mreply := mtype.Out(0) // reply must be an error not anything else

	// 	srv.methods[mname] = &methodType{
	// 		Method:    reflect.TypeOf(handlerStruct).Method(i),
	// 		ArgType:   mtype,
	// 		ReplyType: mreply,
	// 	}
	// }

	// storefront in serviceMap
	serviceMap.Store(srvName, srv) // Handlers1: service1
}
