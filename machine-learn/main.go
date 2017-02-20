package main

import (
	"fmt"
	. "github.com/chewxy/gorgonia"
	"log"
)

func main() {
	g := NewGraph()
	var z, x, y *Node
	var err error
	// define the expression
	x = NewScalar(g, Float64, WithName("x"))
	y = NewScalar(g, Float64, WithName("y"))
	if z, err = Add(x, y); err != nil {
		log.Fatal(err)
	} else {
		log.Println(z.Value())
	}
	// compile into a program
	prog, locMap, err := Compile(g)
	if err != nil {
		log.Fatal(err)
	}
	// create a VM to run the program on
	machine := NewTapeMachine(prog, locMap)
	// set initial values then run
	Let(x, 10.0)
	Let(y, 2.5)

	if err = machine.RunAll(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", z.Value())
}
