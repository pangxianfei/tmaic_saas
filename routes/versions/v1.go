package versions

import (
	"github.com/foolin/gin-template"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/route"
	"net/http"
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

	//engine.LoadHTMLGlob("./resources/views/**/*")
	//处理静态资源（不建议gin框架处理静态资源，参见 public/readme.md 说明 ）
	engine.Static("/public", "./public")
	engine.StaticFS("/dir", http.Dir("./public"))
	engine.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "./resources/views",
		Extension: ".html",
		Master:    "layouts/master",

		DisableCache: true,
	})

	//用户中心
	ver.Auth("", func(grp route.Grouper) {
		//grp.AddGroup("/user", &groups.UserGroup{})
	})

	ver.NoAuth("", func(grp route.Grouper) {
		grp.AddGroup("", &groups.UserAffiliationGroup{})
	})
}
