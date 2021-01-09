package controllers

import (
	"gitee.com/pangxianfei/frame/helpers/tmaic"
	"gitee.com/pangxianfei/frame/http/controller"
	"gitee.com/pangxianfei/frame/request"
	"gitee.com/pangxianfei/frame/saas"

	"gorm.io/gorm"

	"tmaic/app/models"
)

type dbs struct {
	Db   *gorm.DB
	Name string
}

type Saas struct {
	controller.BaseController
}

func (sa *Saas) Index(c request.Context) {
	// db, _ := saas.SetDb(c)
	var id int64
	var TenantId int64
	TenantId = c.GetInt64("TenantId")
	id = c.GetInt64("userid")
	/* 上下文连接db
	user := &models.User{
		ID: id,
	}
	db.Select("user_name", "tenants_id","user_email").First(&user).Scan(&user)
	*/

	// 租户id连接db
	user2 := &models.User{
		ID: id,
	}
	db2, _ := saas.SetTenantDB(TenantId)
	db2.Select("user_name", "tenants_id", "user_email").First(&user2).Scan(&user2)

	// cache.Add("test","pangolin",zone.Now())
	// tmaic.ReturnJson(c,200,200,"yes",user)

	tmaic.Success(c, "yes", user2)

}
