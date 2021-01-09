package versions

import (
	"gitee.com/pangxianfei/frame/config"
	"gitee.com/pangxianfei/frame/request"
	"gitee.com/pangxianfei/frame/route"

	"tmaic/app/http/middleware"
	"tmaic/routes/groups"
)

// RouteV1 版本1路由
func RouteV1(engine *request.Engine) {
	ver := route.NewVersion(engine, "v1")

	ver.Auth(config.GetString("auth.sign_key"), "", func(group route.Grouper) {
		group.AddGroup("/user", &groups.UserGroup{})
		group.AddGroup("/saas", &groups.AaasGroup{})
		group.AddGroup("/user-affiliation", &groups.UserAffiliationGroup{})
	}, middleware.TenantDbMiddle())

	ver.NoAuth("", func(group route.Grouper) {
		group.AddGroup("", &groups.AuthGroup{})
		group.AddGroup("", &groups.BtGroup{})
	})
}
