package main

import "fmt"

func main() {
	// var h Handlers1
	Register(Handlers1{})

	srvName := "Handlers1"
	mName := "Foo"

	service, ok := serviceMap.Load(srvName)
	if !ok {
		fmt.Println("It is not ok for loading the ", srvName)
	}

	method := service.methods[nName]
}
