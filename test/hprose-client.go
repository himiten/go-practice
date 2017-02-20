package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
	"time"
)

type Person struct {
	Age int `age:"18"`
	Name string `name:"scott"`
	Birth time.Time

}

type Stub struct {
	Hello func(string) (string, error) `simple:"true"`
	AsyncHello func(func(string, error), string) `name:"hello"`
	GetPerson func() (*Person,error)
	GenMatrix func(int,int)([][]float64,error)
}


func main(){

	client := rpc.NewClient("http://127.0.0.1:5000/hello")
	client1:= rpc.NewClient("http://127.0.0.1:5000/genMatrix")
	client2:= rpc.NewClient("http://127.0.0.1:5000/getPerson")
	var stub *Stub
	client.UseService(&stub)
	client1.UseService(&stub)
	client2.UseService(&stub)
	defer client.Close()
	defer client1.Close()
	defer client2.Close()
	p,_:=stub.GetPerson()
	mt,_:=stub.GenMatrix(3,3)
	fmt.Printf("%T\n%v\n",mt,mt)
	fmt.Printf("%T\n%v\n",p,p)

	fmt.Println(stub.Hello("World"))



}
