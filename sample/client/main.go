package main

import (
	"context"
	"fmt"
	"rpc/client"
	"rpc/sample/shared"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	// for i := 0; i < 100; i++ {
	wg.Add(1)
	go func() {
		fmt.Println("=========")

		c, err := client.NewRPCClient("tcp", ":8888", client.DefaultOption)
		if err != nil {
			panic(err)
		}
		fmt.Println("=========")

		args := shared.Args{
			A: 200,
			B: 300,
		}

		reply := &shared.Reply{}
		fmt.Println("=========Call")
		if err = c.Call(context.TODO(), "Arith.XX", args, reply); err != nil {
			fmt.Println("=========1")
			panic(err)
		} else {
			fmt.Println("=========")
			fmt.Println(reply.C)
		}

		wg.Done()
	}()

	wg.Wait()
	// }
}
