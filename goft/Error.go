package goft

import "github.com/gin-gonic/gin"

func Error(err error, msg ...string) {
	if err == nil {
		return
	}
	errMsg := err.Error()
	if len(msg) > 0 {
		errMsg = msg[0]
	}
	panic(errMsg)
}

// todo 这里需要抽象成接口外部来操作
func ErrorFunc() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				context.AbortWithStatusJSON(400, gin.H{
					"error": err,
				})
			}
		}()
		context.Next()
	}
}
