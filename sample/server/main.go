package main

import (
	"fmt"
	"rpc/server"
	"time"
)

func main() {
	s := server.NewSimpleServer(server.DefaultOption)
	err := s.Register(Arith{}, make(map[string]string))
	if err != nil {
		panic(err)
	}

	go func() {
		err = s.Serve("tcp", ":8888")
		if err != nil {
			fmt.Println("//////ERR: ", err)
			panic(err)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
