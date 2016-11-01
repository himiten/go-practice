package main
import (
	"fmt"
	"flag"
)
var address = flag.String("addr", ":1883", "listen address of broker")
func main() {
	flag.Parse()
	fmt.Printf("%T\n",address)
	fmt.Printf("%v\n",*address)
}