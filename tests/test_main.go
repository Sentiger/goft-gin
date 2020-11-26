package main

import (
	"github.com/sentiger/goft-gin/goft"
	. "github.com/sentiger/goft-gin/tests/controllers"
	. "github.com/sentiger/goft-gin/tests/exceptions"
	. "github.com/sentiger/goft-gin/tests/middlewares"
	"log"
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
		Task("0/3 * * * * *", func() {
			log.Println("我是定时任务1")
		}).
		Task("0/1 * * * * *", func() {
			log.Println("我是定时任务2")
		}).
		Launch()
}
