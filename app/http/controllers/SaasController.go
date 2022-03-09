package controllers

import (
	"gitee.com/pangxianfei/frame/helpers/tmaic"
	"gitee.com/pangxianfei/frame/http/controller"
	"gitee.com/pangxianfei/frame/request"
	"gitee.com/pangxianfei/frame/saas"
	"tmaic/app/http/requests"
	"tmaic/app/saasmodels"
)

type Saas struct {
	controller.BaseController
}

func (sa *Saas) Index(c request.Context) {

	var requestData requests.TenantDBRegister
	if !sa.ValidateJSON(c, &requestData, true) {
		return
	}

	//var TenantId int64
	//userid := c.GetInt64("userid")

	//fmt.Println("\r\n==========================\r\n")
	//fmt.Println(requestData.TenantId)
	//fmt.Println("\r\n****************\r\n")
	user := &saasmodels.User{}
	//var TenantUser saasmodels.User
	//通过上下文连接数据库
	//DBS := TenantUser.DB(c)
	//DBS.First(&user).Debug()
	//通过租户ID连接数据库
	DBS2 := saas.SetTenantDB(requestData.TenantId)
	DBS2.First(&user).Debug()
	//dbid := c.PostForm("dbid")
	//fmt.Println(dbid)
	tmaic.Success(c, "yes", user)
}
