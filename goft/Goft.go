package goft

import "github.com/gin-gonic/gin"

type Goft struct {
	*gin.Engine
}

func Ignite() *Goft {
	g := &Goft{
		Engine: gin.New(),
	}
	return g
}

func (goft *Goft) Launch() {
	goft.Run(":8081")
}
