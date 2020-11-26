package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	g         *gin.RouterGroup
	registers []interface{}
}

// 初始化
func Ignite() *Goft {
	g := &Goft{
		Engine:    gin.New(),
		registers: make([]interface{}, 1),
	}
	return g
}

// 注册类
func (goft *Goft) register(class interface{}) {
	goft.registers = append(goft.registers, class)
}

func (goft *Goft) RegisterExceptionHandler(exceptionHandler IExceptionHandler) *Goft {
	defaultExceptionHandler := NewExceptionHandler()
	if exceptionHandler == nil {
		exceptionHandler = defaultExceptionHandler
	}
	goft.Use(defaultExceptionHandler.HandlerFunc(exceptionHandler))
	return goft
}

// 重写gin.Handle方法，这里可以直接加载group，省略在控制器中还要加这个参数
func (goft *Goft) Handle(httpMethod, relativePath string, handler interface{}) *Goft {
	//if h,ok:=handler.(func(ctx *gin.Context)string);ok {
	//	goft.g.Handle(httpMethod, relativePath, func(context *gin.Context) {
	//		context.String(200,h(context))
	//	})
	//}
	if h := Convert(handler); h != nil {
		goft.g.Handle(httpMethod, relativePath, h)
	}
	return goft
}

// 中间件
func (goft *Goft) Middleware(mid IMiddleware) *Goft {
	goft.Use(func(context *gin.Context) {
		// 其实这里可以不进行处理，直接就是mid.OnRequest(context)，在具体的中间件中进行特殊处理。
		err := mid.OnRequest(context)
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"err": err.Error()})
		} else {
			context.Next()
		}
	})
	return goft
}

// 挂在控制器
func (goft *Goft) Mount(group string, controllers ...IController) *Goft {
	goft.g = goft.Group(group)
	for _, controller := range controllers {
		controller.Build(goft)
	}
	return goft
}

// 启动程序
func (goft *Goft) Launch() {
	config := InitConfig()
	goft.Run(fmt.Sprintf(":%d", config.Server.Port))
}
