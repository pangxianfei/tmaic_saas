package versions

import (
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/route"
	"tmaic/routes/groups"
)

//后端路由
func V1Api(engine *request.Engine) {
	ver := route.NewVersion(engine, "v1")

	ver.Auth("", func(grp route.Grouper) {
		grp.AddGroup("/user", &groups.UserGroup{})
	})

	ver.NoAuth("", func(grp route.Grouper) {
		grp.AddGroup("", &groups.AuthGroup{})
		grp.AddGroup("", &groups.ArticleGroup{})
	})
}

//前端路由
func WebRoute(engine *request.Engine) {
	ver := route.NewVersion(engine, "")

	//用户中心
	ver.Auth("", func(grp route.Grouper) {
		//grp.AddGroup("/user", &groups.UserGroup{})
	})

	ver.NoAuth("", func(grp route.Grouper) {
		//grp.AddGroup("", &groups.AuthGroup{})
		//grp.AddGroup("", &groups.ArticleGroup{})
		grp.AddGroup("", &groups.UserAffiliationGroup{})
	})
}
