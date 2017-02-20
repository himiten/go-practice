package main

import (
  "gopkg.in/kataras/iris.v6"
  "gopkg.in/kataras/iris.v6/adaptors/httprouter"
  "fmt"
)
type User struct {
 Id string
 Name string
}
func main(){
  app := iris.New()
  app.Adapt(iris.DevLogger())
  app.Adapt(httprouter.New())
  myMiddleware := func(ctx *iris.Context) {
    println(ctx.Method() + ": " + ctx.Path())
    ctx.Next()
  }
  group:=app.Party("/hello",myMiddleware)
  {
    group.Get("/:id/:name",func(ctx *iris.Context){
      id:=ctx.Param("id")
      name:=ctx.Param("name")
      age:=ctx.URLParam("age")
      fmt.Printf("type:%T,id=%v\n",id,id)
      fmt.Printf("type:%T,name=%v\n",name,name)
      fmt.Printf("type:%T,age=%v\n",age,age)
      ctx.JSON(iris.StatusOK,User{id,name})

    })
  }
  app.Listen("127.0.0.1:6001")
}
