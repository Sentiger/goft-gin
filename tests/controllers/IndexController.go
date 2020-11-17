package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sentiger/goft-gin/goft"
	"github.com/sentiger/goft-gin/tests/models"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (this *IndexController) GetIndex(ctx *gin.Context) string {
	return "Hello goft-gin"
}

func (this *IndexController) InfoIndex(ctx *gin.Context) goft.IModel {
	panic("强行异常")
	return &models.UserModel{
		Uid:      1,
		Username: "张三",
	}
}

func (this *IndexController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/home", this.GetIndex)
	goft.Handle("GET", "/info", this.InfoIndex)
}
