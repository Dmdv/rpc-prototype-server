package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"rpc-prototype-server/pkg"
)

var (
	addr = flag.String("addr", ":8972", "server address")
)

type Arith struct{}

// Mul the second parameter is not a pointer
func (t *Arith) Mul(_ context.Context, args pkg.Args, reply *pkg.Reply) error {
	reply.C = args.A * args.B
	fmt.Println("C=", reply.C)
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	err := s.Register(new(Arith), "")
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server")
	err = s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
