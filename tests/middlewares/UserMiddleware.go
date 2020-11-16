package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {

}

func (this *UserMiddleware)OnRequest(ctx *gin.Context)error  {
	fmt.Println("用户中间件")
	//return fmt.Errorf("错误")
	ctx.Next()
	return nil
}

func NewUserMiddleware()*UserMiddleware  {
	return &UserMiddleware{}
}
