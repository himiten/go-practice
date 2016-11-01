package main

import (
	"github.com/go-martini/martini"
	"log"
	//"os"
)

func main() {
	m := martini.Classic()

	m.Use(martini.Static("public"))

	m.Handlers(
		func(c martini.Context,log *log.Logger) {
			log.Println("before a request for md1")
			c.Next()
			log.Println("after a request for md1")
		},
		func(c martini.Context,log *log.Logger) {
			log.Println("before a request for md2")
			c.Next()
			log.Println("after a request for md2")
		},
	)
	m.Get("/", func() string {
		return "Hello goweb!"
	})

	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})
	m.Run()
}
