package controllers

import (
	"github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/request"
)

type UserAffiliation struct {
	controller.BaseController
}

func (this *UserAffiliation) Index(c request.Context) {

	this.TplName = "index"
	this.Data = tmaic.Output{"Content": "duzhenxun2", "colors": "colors", "red": "red"}
	//article := models.Article{}
	//result := db.First(&article)
	//c.HTML(http.StatusOK, this.TplName, result) gin 原生写法
	this.View(c)
}
