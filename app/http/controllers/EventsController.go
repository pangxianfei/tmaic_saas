package controllers

import (
	"gitee.com/pangxianfei/frame/helpers/log"
	"gitee.com/pangxianfei/frame/helpers/tmaic"
	"gitee.com/pangxianfei/frame/http/controller"
	"gitee.com/pangxianfei/frame/hub"
	"gitee.com/pangxianfei/frame/request"

	"tmaic/app/events"
	pbs "tmaic/app/events/protocol_model"
)

type Events struct {
	controller.BaseController
}

func (Events *Events) EventsDemo(c request.Context) {

	ur := events.Test{}
	param := &pbs.Test{
		Id: 1,
	}

	ur.SetParam(param)
	if errs := hub.Emit(&ur); errs != nil {
		log.Info("user registered event emit failed", tmaic.V{"event": ur, "errors": errs})
	}

	tmaic.Success(c, "yes", param)
}
