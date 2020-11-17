package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sentiger/goft-gin/goft"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (this *IndexController) GetIndex(ctx *gin.Context) string {
	return "Hello goft-gin"
}

func (this *IndexController) InfoIndex(ctx *gin.Context) string {
	return "Hello goft-gin info"
}

func (this *IndexController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/home", this.GetIndex)
	goft.Handle("GET", "/info", this.InfoIndex)
}
