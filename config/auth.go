package config

import (
	. "github.com/pangxianfei/framework/config"
	"tmaic/app/models"
)

func init() {
	auth := make(map[string]interface{})

	auth["sign_key"] = Env("AUTH_SIGN_KEY", "sign key")
	//必须是指针
	auth["model_ptr"] = &models.User{} 

	Add("auth", auth)
}
