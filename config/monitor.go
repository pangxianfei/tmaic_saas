package config

import (
	. "github.com/pangxianfei/framework/config"
)

func init() {
	tmaic := make(map[string]interface{})
	tmaic["port"] = Env("MONITOR_PORT", "8080")
	tmaic["username"] = Env("MONITOR_USERNAME", "tmaic")
	tmaic["password"] = Env("MONITOR_PASSWORD", "Password")

	Add("tmaic", tmaic)
}
