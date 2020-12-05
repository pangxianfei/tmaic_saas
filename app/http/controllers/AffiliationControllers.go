package controllers

import (
	tmaic "github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/request"
)

type UserAffiliation struct {
	controller.BaseController
}

//var template = new(view.Template) interface{}

func init() {

	log.Debug("UserAffiliation")

}

func (this *UserAffiliation) Index(c request.Context) {

	this.TplName = "nodes"

	this.Data = tmaic.Output{"Content": "duzhenxun2", "color": "red"}

	this.View(c)
	//以下方法一样可以调用模板渲染
	//c.View("nodes",tmaic.V{"Content": "pangxianfei",})

}
