package goft

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

var ResponderList []IResponder

func init() {
	ResponderList = []IResponder{
		new(StringResponder),
		new(ModelResponder),
	}
}

type IResponder interface {
	RespondTo() gin.HandlerFunc
}

func Convert(handler interface{}) gin.HandlerFunc {
	hRef := reflect.ValueOf(handler)
	hType := hRef.Type()
	for _, r := range ResponderList {
		rRef := reflect.ValueOf(r).Elem()
		if hType.ConvertibleTo(rRef.Type()) {
			rRef.Set(hRef)
			return rRef.Interface().(IResponder).RespondTo()
		}
	}
	return nil
}

// 定义string的返回类型处理
type StringResponder func(*gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, this(context))
	}
}

// 返回一个模型
type ModelResponder func(*gin.Context) IModel

func (this ModelResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, this(context))
	}
}
