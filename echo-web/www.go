package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/hprose/hprose-golang/rpc"
	alg "github.com/himiten/go-practice/algorithm-practice/algorithm"

	"fmt"
	//"strconv"
	"time"
)
func hello(name string) string {
	return "Hello " + name + "!"
}

type Person struct {
	Age int `age:"18"`
	Name string `name:"scott"`
	Birth time.Time
}
func genMatrix(rows,cols int) [][]float64{
	temp:=alg.GenMatrix(rows,cols)
	fmt.Println(temp)
	/*const ROW  = rows
	const COL  = cols
*/
	return temp
	/*var str string
	for i:=0;i<rows;i++ {
		for j:=0;j<cols;j++{
			str+=strconv.FormatFloat(temp[i][j],'e',6,64)
		}
	}

	fmt.Println(str)
	return str*/
}
func getPerson() *Person {
	return &Person{18,"world",time.Now()}
}


func main() {
	e := echo.New()
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	service.AddFunction("genMatrix", genMatrix)
	service.AddFunction("getPerson", getPerson)

	e.Any("/hello", echo.WrapHandler(service))
	e.Any("/genMatrix", echo.WrapHandler(service))
	e.Any("/getPerson", echo.WrapHandler(service))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":5001"))
}
