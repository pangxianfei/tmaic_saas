package config

import (
	. "github.com/pangxianfei/framework/config"
)

func init() {
	app := make(map[string]interface{})

	app["name"] = Env("APP_NAME", "tmaic")
	app["env"] = Env("APP_ENV", "production")
	app["debug"] = Env("APP_DEBUG", false)
	app["log_level"] = Env("APP_LOG_LEVEL", "trace")
	app["port"] = Env("APP_PORT", "80")
	app["timezone"] = "UTC"
	app["locale"] = Env("APP_LOCALE", "en")
	app["fallback_locale"] = "en"
	app["key"] = Env("APP_KEY")
	app["cipher"] = "AES-256-CBC"
	app["read_timeout_seconds"] = 10
	app["write_timeout_seconds"] = 10

	Add("app", app)
}
