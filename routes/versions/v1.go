package versions

import (
	"github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/route"
	"tmaic/routes/groups"
)

func NewV1(engine *request.Engine) {
	ver := route.NewVersion(engine, "v1")

	// auth routes
	ver.Auth(config.GetString("auth.sign_key"), "", func(grp route.Grouper) {
		grp.AddGroup("/user", &groups.UserGroup{})
	})

	// no auth routes
	ver.NoAuth("", func(grp route.Grouper) {
		grp.AddGroup("", &groups.AuthGroup{})
		//grp.AddGroup("/user-affiliation", &groups.UserAffiliationGroup{})
	})
}
