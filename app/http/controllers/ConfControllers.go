package controllers

import (
	"github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/request"
	"net/http"
)

type Conf struct {
	controller.BaseController
}

func (c *Conf) Index(request request.Context) {

	request.JSON(http.StatusOK, tmaic.Output{"data": "user"})
	return
}
