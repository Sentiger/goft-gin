package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type IExceptionHandler interface {
	Render(*gin.Context, interface{})
}

type ExceptionHandler struct {
}

func NewExceptionHandler() *ExceptionHandler {
	return &ExceptionHandler{}
}

func (this *ExceptionHandler) Render(ctx *gin.Context, err interface{}) {
	fmt.Println("我是系统异常")
}

func (this *ExceptionHandler) HandlerFunc(handler IExceptionHandler) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			e := recover()
			if e != nil {
				handler.Render(context, e)
			}
		}()
		context.Next()
	}
}
