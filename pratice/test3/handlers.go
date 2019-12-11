package main

import (
	"context"
	"fmt"
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

// Handlers1 ..
type Handlers1 struct{}

// Foo ..
func (h Handlers1) Foo(ctx context.Context, arg *Args, reply *Reply) error {
	fmt.Println("Foo is calling with", arg.A, arg.B)

	reply.r = "Hello Norman, this is the Foo result"

	return nil
}

// Goo ..
func (h Handlers1) Goo(ctx context.Context, arg Args, reply *Reply) error {
	fmt.Println("Goo is calling with", arg.A, arg.B)

	reply.r = "Hello Norman, this is the Goo result"

	return nil
}

// Zoo ..
func (h *Handlers1) Zoo(ctx context.Context, arg Args, reply *Reply) error {
	fmt.Println("Zoo is calling with", arg.A, arg.B)

	reply.r = "Hello Norman, this is the Zoo result"

	return nil
}
