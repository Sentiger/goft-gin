package goft

import "github.com/gin-gonic/gin"

type IMiddleware interface {
	OnRequest(*gin.Context)error
}
