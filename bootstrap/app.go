package bootstrap

import (
	"fmt"
	"gitee.com/pangxianfei/frame/cache"
	"gitee.com/pangxianfei/frame/database"
	"gitee.com/pangxianfei/frame/helpers/m"
	"gitee.com/pangxianfei/frame/helpers/zone"
	"gitee.com/pangxianfei/frame/http/middleware"
	"gitee.com/pangxianfei/frame/logs"
	//"gitee.com/pangxianfei/frame/queue"
	"gitee.com/pangxianfei/frame/request"
	"gitee.com/pangxianfei/frame/sentry"
	"gitee.com/pangxianfei/frame/validator"

	"tmaic/config"
	"tmaic/resources/lang"

	c "gitee.com/pangxianfei/frame/config"
)

func Initialize() {
	config.Initialize()
	sentry.Initialize()
	logs.Initialize()
	zone.Initialize()
	lang.Initialize()
	cache.Initialize()
	database.Initialize()
	m.Initialize()

	//queue.Initialize()
	//jobs.Initialize()
	//events.Initialize()
	//listeners.Initialize()

	validator.UpgradeValidatorV8toV9()
}

func Middleware(r *request.Engine) {
	r.Use(middleware.RequestLogger())

	if c.GetString("app.env") == "production" {
		r.Use(middleware.Logger())
		r.Use(middleware.Recovery())
	}
	fmt.Println("中间件")
	r.Use(middleware.Locale())
}
