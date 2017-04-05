package main

import (
  "gopkg.in/kataras/iris.v6"
  "gopkg.in/kataras/iris.v6/adaptors/httprouter"
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
)
type User struct {
  Id string
  Name string
}
func main(){
  app := iris.New()
  app.Adapt(iris.DevLogger())
  app.Adapt(httprouter.New())

  app.Get("/img/:imgName",func(ctx *iris.Context){
    res := ctx.ResponseWriter
    imgName :=ctx.Param("imgName")
    imgDir,_ := filepath.Abs("./iris-web/assets/img")
    imgPath := filepath.Join(imgDir,imgName+".png")
    fmt.Println(imgPath)
    buf,err:=ioutil.ReadFile(imgPath)
    if err!=nil{
      fmt.Println(err)
      fmt.Fprintf(os.Stderr,"File Error:%s\n",err)
      res.WriteString("文件读取错误！")
      res.Flush()
    }
    res.SetContentType("image/png ")
    res.Write(buf)
    res.Flush()
  })


  myMiddleware := func(ctx *iris.Context) {
    println(ctx.Method() + ": " + ctx.Path())
    ctx.Next()
  }

  group:=app.Party("/hello",myMiddleware)
  {
    group.Get("/",func(ctx *iris.Context){
      ctx.ResponseWriter.WriteString("hello iris")
      ctx.ResponseWriter.Flush()
    })

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
  app.Listen(":8080")
}
