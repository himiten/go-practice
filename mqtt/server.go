package main

import (

	"flag"
	"log"
	"net"

	"github.com/jeffallen/mqtt"
)
var addr = flag.String("addr", ":1883", "listen address of broker")
func main() {
	flag.Parse()
	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen errror: %#v", err)
	}
	svr := mqtt.NewServer(l)
	svr.Start()
	<-svr.Done
}
