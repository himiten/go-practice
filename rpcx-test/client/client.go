package main

import (
	"time"
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/log"
)

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

func main() {
	s := &rpcx.DirectClientSelector{Network: "tcp", Address: "127.0.0.1:8972",DialTimeout: 10 * time.Second}
	client := rpcx.NewClient(s)
	a,b:=Complex{3,2},Complex{3,-2}
	args := &Args{a, b}
	var reply Reply
	err := client.Call("complex.Mul", args, &reply)
	if err != nil {
		log.Infof("error for Complex: %v*%v, %v \n", args.A, args.B, err)
	} else {
		log.Infof("Complex: %v*%v=%v \n", args.A, args.B, reply.C)
	}

	client.Close()
}