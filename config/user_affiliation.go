package config

import (
	. "github.com/pangxianfei/framework/config"
)

func init() {
	userAffiliation := make(map[string]interface{})

	userAffiliation["enable"] = true

	Add("user_affiliation", userAffiliation)
}
