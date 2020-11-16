package main

import (
	"github.com/sentiger/goft-gin/goft"
	. "github.com/sentiger/goft-gin/tests/controllers"
	. "github.com/sentiger/goft-gin/tests/middlewares"
)

func main() {
	goft.Ignite().
		Middleware(NewUserMiddleware()).
		Mount("v1", NewIndexController()).
		Mount("v2", NewIndexController()).
		Launch()

	//g := gin.New()
	//r := g.Group()
	//r.Handle()
}
