package routes

import (
	"gitee.com/pangxianfei/frame/request"
	"gitee.com/pangxianfei/frame/route"

	"tmaic/routes/versions"
)

func Register(router *request.Engine) {
	defer route.Bind(router)

	versions.RouteV1(router)
}
