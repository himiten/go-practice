package main

import (
	//"runtime"

	"github.com/hprose/hprose-golang/rpc"
	"github.com/kataras/iris"
	//"github.com/valyala/fasthttp"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	/*runtime.GOMAXPROCS(runtime.NumCPU())
	service := rpc.NewFastHTTPService()
	service.AddFunction("greet", hello)
	fasthttp.ListenAndServe(":8080", service.ServeFastHTTP)*/

	service := rpc.NewFastHTTPService()
	service.AddFunction("hello", hello)
	iris.Any("/hello", func(c *iris.Context) {
		service.ServeFastHTTP(c.RequestCtx)
	})
	iris.Listen(":8080")
}