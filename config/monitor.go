package config

import (
	. "gitee.com/pangxianfei/frame/config"
)

func init() {
	monitor := make(map[string]interface{})

	monitor["port"] = Env("MONITOR_PORT", "8080")
	monitor["username"] = Env("MONITOR_USERNAME", "Totoval")
	monitor["password"] = Env("MONITOR_PASSWORD", "Passw0rd")

	Add("monitor", monitor)
}
