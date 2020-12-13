package config

import (
	. "github.com/pangxianfei/framework/config"
	"tmaic/app/models"
)

func init() {
	auth := make(map[string]interface{})

	auth["sign_key"] = Env("AUTH_SIGN_KEY", "8ka5Ekes9Eel65ksf6E4rnLJ6pPmJ5f5SP146W")
	//必须是指针
	auth["model_ptr"] = &models.User{}

	Add("auth", auth)
}
