package config

import (
	. "gitee.com/pangxianfei/frame/config"
)

func init() {
	userAffiliation := make(map[string]interface{})

	userAffiliation["enable"] = true

	Add("user_affiliation", userAffiliation)
}
