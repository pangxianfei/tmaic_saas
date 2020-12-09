package config

import (
	. "github.com/pangxianfei/framework/config"
)

func init() {
	app := make(map[string]interface{})

	app["name"] = Env("APP_NAME", "tmaic")
	app["env"] = Env("APP_ENV", "production")
	app["debug"] = Env("APP_DEBUG", false)
	app["app_log"] = Env("APP_LOG", false)
	app["log_level"] = Env("APP_LOG_LEVEL", "trace")
	app["port"] = Env("APP_PORT", "80")
	app["timezone"] = "UTC"
	app["locale"] = Env("APP_LOCALE", "en")
	app["fallback_locale"] = "zh"
	app["key"] = Env("APP_KEY")
	app["cipher"] = "AES-256-CBC"
	app["read_timeout_seconds"] = 10
	app["write_timeout_seconds"] = 10

	app["public_key"] = Env("PUBLIC_KEY", "./storage/cert/public.pem")
	app["private_key"] = Env("PRIVATE_KEY", "./storage/cert/private.pem")

	app["TEMPLATE_TPL"] = Env("TEMPLATE_TPL", "./resources/views/")
	app["SUFFIX"] = Env("SUFFIX", ".html")
	app["layout"] = Env("LAYOUTS", ".html")
	app["public_path"] = Env("PUBLIC_PATH", "/public")

	//token
	app["token_encryption"] = Env("TOKE_KEY_ENCRYPTION", false)

	Add("app", app)
}
