package config

import (
	. "gitee.com/pangxianfei/frame/config"
	"tmaic/app/models"
)

func init() {
	auth := make(map[string]interface{})

	auth["sign_key"] = Env("AUTH_SIGN_KEY", "sign key")
	auth["model_ptr"] = &models.User{} // must be a pointer

	Add("auth", auth)
}
