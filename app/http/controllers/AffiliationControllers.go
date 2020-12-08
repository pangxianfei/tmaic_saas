package controllers

import (
	tmaic "github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/request"
	"log"
)

type Affiliation struct {
	controller.BaseController
}

func (this *Affiliation) All(c request.Context) {

	this.ShowData = tmaic.Output{
		"title":       "GoLang API Framework 开发框架",
		"keywords":    "Golang开发框架,Golang web mvc框架,GoLang API Framework 开发框架,gin框架,mvc设计模式",
		"description": "tmaic 是一套简洁、优雅的Golang API Web开发框架(GoLang Web Framework)。支持mysql,mssql等多类型数据库，它可以让你从面条一样杂乱的代码中解脱出来；它可以帮你构建一个完美的网络应用，而且每行代码都可以简洁、富于表达力。",
	}
	/*
		var pc uintptr
		var file string
		var line int
		var ok bool
		pc, file, line, ok = runtime.Caller(0)
		log.Println(pc)
		log.Println(file)
		log.Println(line)
		log.Println(ok)
		var f *runtime.Func
		f = runtime.FuncForPC(pc)
		log.Println(f.Name())
	*/
	log.Println(c.HandlerNames())

	this.View(c)

}
