package controllers

import (
	"fmt"
	"gitee.com/pangxianfei/frame/helpers/tmaic"
	"gitee.com/pangxianfei/frame/http/controller"
	"gitee.com/pangxianfei/frame/request"
	"gitee.com/pangxianfei/frame/saas"
	"tmaic/app/models"

	"tmaic/app/saasmodels"
)

type Saas struct {
	controller.BaseController
}

func (sa *Saas) Index(c request.Context) {
	var id int64
	//var TenantId int64

	user := &models.User{
		ID: id,
	}

	user3 := &models.User{
		ID: id,
	}
	//TenantId = c.GetInt64("TenantId")
	//db2, _ := saas.SetTenantDB(TenantId)
	//db2.Select("user_name", "tenants_id", "user_email").First(&user).Scan(&TenantUser)

	id = c.GetInt64("userid")
	var TenantUser saasmodels.SaasUser
	var TenantUser3 saasmodels.SaasUser

	DBS := TenantUser.DB(c)

	fmt.Println(DBS.Name())

	DBS.First(&user).Debug()

	db3 := saas.SetTenantDB(1)
	db3.Select("user_name", "tenants_id", "user_email").First(&user3).Scan(&TenantUser3)

	fmt.Println(TenantUser3.Email)
	tmaic.Success(c, "yes", user)
}
