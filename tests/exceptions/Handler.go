package exceptions

import (
	"github.com/gin-gonic/gin"
	"github.com/sentiger/goft-gin/goft"
)

type Handler struct {
	goft.ExceptionHandler
}

func NewHandler() *Handler {
	return &Handler{}
}

func (this *Handler) Render(ctx *gin.Context, err interface{}) {
	ctx.AbortWithStatusJSON(400, gin.H{"error": err})
}
