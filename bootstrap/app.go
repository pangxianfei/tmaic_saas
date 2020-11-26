package bootstrap

import (
	"github.com/pangxianfei/framework/cache"
	"github.com/pangxianfei/framework/database"
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/helpers/zone"
	"github.com/pangxianfei/framework/http/middleware"
	"github.com/pangxianfei/framework/logs"
 
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/sentry"
	"github.com/pangxianfei/framework/validator"

 
	"tmaic/config"
	"tmaic/resources/lang"

	c "github.com/pangxianfei/framework/config"
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
	 

	validator.UpgradeValidatorV8toV9()
}

func Middleware(r *request.Engine) {
	r.Use(middleware.RequestLogger())

	if c.GetString("app.env") == "production" {
		r.Use(middleware.Logger())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.Locale())
}
