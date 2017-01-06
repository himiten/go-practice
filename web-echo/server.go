package main

import (
	//"net/http"
	//"io/ioutil"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Static("/", "static")

	e.File("/", "index.html")

	/*e.GET("/", func(c echo.Context) error {
		buf,err := ioutil.ReadFile("./index.html")
		if err!=nil{
			e.Logger.Error("读取html错误",err)
		}else {
			e.Logger.Info("正确获取html")
		}
		return c.HTMLBlob(http.StatusOK, buf)
	})*/
	e.Logger.Fatal(e.Start(":3001"))
}
