package main

import (
	"context"
	"errors"
	"fmt"
	"rpc/sample/shared"
)

type Arith struct{}

//arg可以是指针类型，也可以是指针类型
func (a Arith) Add(ctx context.Context, arg *shared.Args, reply *shared.Reply) error {
	fmt.Println("=====In Add")

	reply.C = arg.A + arg.B
	return nil
}

func (a Arith) Minus(ctx context.Context, arg shared.Args, reply *shared.Reply) error {
	reply.C = arg.A - arg.B
	return nil
}

func (a Arith) Mul(ctx context.Context, arg shared.Args, reply *shared.Reply) error {
	reply.C = arg.A * arg.B
	return nil
}

func (a Arith) Divide(ctx context.Context, arg *shared.Args, reply *shared.Reply) error {
	if arg.B == 0 {
		return errors.New("divided by 0")
	}

	reply.C = arg.A / arg.B
	return nil
}
