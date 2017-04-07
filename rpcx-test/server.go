package main

import "github.com/smallnest/rpcx"


type Complex struct {
	Real float64
	Image float64
}
type Args struct {
	A Complex `msg:"a"`
	B Complex `msg:"b"`
}

type Reply struct {
	C Complex `msg:"c"`
}

func (t *Complex) Mul(args *Args, reply *Reply) error {
	a,b:=args.A,args.B
	reply.C.Real = a.Real*b.Real-a.Image*b.Image
	reply.C.Image =a.Image*b.Real+a.Real*b.Image
	return nil
}

func (t *Complex) Error(args *Args, reply *Reply) error {
	panic("ERROR")
}

func main() {
	server := rpcx.NewServer()
	server.RegisterName("complex", new(Complex))
	server.Serve("tcp", "127.0.0.1:8972")
}