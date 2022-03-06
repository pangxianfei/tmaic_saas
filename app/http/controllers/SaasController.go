package controllers

import (
	"gitee.com/pangxianfei/frame/helpers/tmaic"
	"gitee.com/pangxianfei/frame/http/controller"
	"gitee.com/pangxianfei/frame/request"
	"tmaic/app/saasmodels"
)

type Saas struct {
	controller.BaseController
}

func (sa *Saas) Index(c request.Context) {

	//var TenantId int64
	userid := c.GetInt64("userid")

	user := &saasmodels.User{
		ID: userid,
	}

	var TenantUser saasmodels.User

	DBS := TenantUser.DB(c)

	DBS.First(&user).Debug()

	tmaic.Success(c, "yes", user)
}
