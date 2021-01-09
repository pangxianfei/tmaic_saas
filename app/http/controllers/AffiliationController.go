package controllers

import (
	"net/http"

	"tmaic/app/models"

	"gitee.com/pangxianfei/frame/helpers/tmaic"

	"gitee.com/pangxianfei/frame/http/controller"
	"gitee.com/pangxianfei/frame/request"
)

type UserAffiliation struct {
	controller.BaseController
}

func (uaff *UserAffiliation) RenderAll(c request.Context) {
	var u models.UserAffiliation
	c.HTML(http.StatusOK, "user_affiliation.nodes", tmaic.V{
		"data": u.All(),
	})

	return
}
