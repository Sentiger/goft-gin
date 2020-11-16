package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sentiger/goft-gin/goft"
	"time"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (this *IndexController) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		time.Sleep(time.Second*5)
		context.JSON(200, gin.H{
			"result": "首页",
		})
	}
}

func (this *IndexController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/home", this.GetIndex())
}
