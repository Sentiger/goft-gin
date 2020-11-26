package main

import (
	"github.com/sentiger/goft-gin/goft"
	. "github.com/sentiger/goft-gin/tests/controllers"
	. "github.com/sentiger/goft-gin/tests/exceptions"
	. "github.com/sentiger/goft-gin/tests/middlewares"
)

func main() {
	//fmt.Println(goft.InitConfig().Server)
	//
	//return
	goft.Ignite().
		RegisterExceptionHandler(NewHandler()). // 注册系统异常
		Middleware(NewUserMiddleware()).        // 注册中间件
		Mount("v1", NewIndexController()).
		Mount("v2", NewIndexController()).
		Middleware(NewUserMiddleware()).
		Launch()
}
