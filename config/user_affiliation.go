package config

import (
	. "github.com/pangxianfei/framework/config"
)

func init() {
	//注入配置 实际运行系统时间注入配置值 通过 GetString("user_affiliation.***") 获取配置值
	user := make(map[string]interface{})

	user["enable"] = true

	Add("user_affiliation", user)
}
