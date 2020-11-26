package controllers

import (
	"github.com/pangxianfei/framework/model"
	"net/http"
	"tmaic/app/models"

	"github.com/pangxianfei/framework/helpers/toto"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/request"
)

type Article struct {
	controller.BaseController
	model.BaseModel
}

func (l *Article) Index(c request.Context) {

	db := l.DB()

	article := models.Article{Title: "这是标题",Body: "这是内容",Slug: 12345}

	db.Create(&article) // 通过数据的指针来创建

	c.JSON(http.StatusUnprocessableEntity, toto.V{"Article": "result"})
	return
}
