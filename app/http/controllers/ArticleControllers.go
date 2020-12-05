package controllers

import (
	"github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/model"
	"net/http"
	"tmaic/app/models"

	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/request"
)

type Article struct {
	controller.BaseController
	model.BaseModel
}

func (l *Article) Index(c request.Context) {

	db := l.DB()

	article := models.Article{Title: "这是标题", Body: "这是内容", Slug: "12345"}
	//db.Migrator().CreateTable(&article)
	// 通过数据的指针来创建
	db.Create(&article)

	//s,_:= tmaic.Encryption(string("Golang使用RSA进行公钥加密私钥解密,私钥加密公钥解密的实现"))

	//dd,_ := tmaic.Decrypt(s)

	c.JSON(http.StatusOK, tmaic.Date("2006-01-02 15:04:05", tmaic.Time()))
	return
}
