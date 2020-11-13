package goft

import "github.com/gin-gonic/gin"

type Goft struct {
	*gin.Engine
}

// 初始化
func Ignite() *Goft {
	g := &Goft{
		Engine: gin.New(),
	}
	return g
}

func (goft *Goft) Mount(controllers ...IController) *Goft {
	for _, controller := range controllers {
		controller.Build(goft)
	}
	return goft
}

// 启动程序
func (goft *Goft) Launch() {

	goft.Run(":8081")
}
