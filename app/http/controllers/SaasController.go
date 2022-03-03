package controllers

import (
	"gitee.com/pangxianfei/frame/helpers/tmaic"
	"gitee.com/pangxianfei/frame/http/controller"
	"gitee.com/pangxianfei/frame/request"
	"gitee.com/pangxianfei/frame/saas"

	"tmaic/app/models"
)

type Saas struct {
	controller.BaseController
}

func (sa *Saas) Index(c request.Context) {
	var id int64
	var TenantId int64
	TenantId = c.GetInt64("TenantId")
	id = c.GetInt64("userid")
	user2 := &models.User{
		ID: id,
	}
	db2, _ := saas.SetTenantDB(TenantId)
	db2.Select("user_name", "tenants_id", "user_email").First(&user2).Scan(&user2)

	tmaic.Success(c, "yes", user2)

}
