package routes

import (
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/route"
	"tmaic/routes/versions"
)

func Register(router *request.Engine) {
	defer route.Bind(router)

	versions.NewV1(router)
}
