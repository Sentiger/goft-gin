package goft

import (
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	g *gin.RouterGroup
}

// 初始化
func Ignite() *Goft {
	g := &Goft{
		Engine: gin.New(),
	}
	return g
}

// 重写gin.Handle方法，这里可以直接加载group，省略在控制器中还要加这个参数
func (goft *Goft) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Goft {
	goft.g.Handle(httpMethod, relativePath, handlers...)
	return goft
}

// 中间件
func (goft *Goft)Middleware(mid IMiddleware) *Goft {
	goft.Use(func(context *gin.Context) {
		// 其实这里可以不进行处理，直接就是mid.OnRequest(context)，在具体的中间件中进行特殊处理。
		err:=mid.OnRequest(context)
		if err!=nil {
			context.AbortWithStatusJSON(400, gin.H{"err":err.Error()})
		}else{
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
	goft.Run(":8081")
}
