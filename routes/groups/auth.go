package groups

import (
	"github.com/pangxianfei/framework/route"
	"tmaic/app/http/controllers"
)

type AuthGroup struct {
	LoginController    controllers.Login
	RegisterController controllers.Register
}

func (ag *AuthGroup) Group(group route.Grouper) {
	group.POST("/login", ag.LoginController.Login).Name("login")
	group.POST("/register", ag.RegisterController.Register)
}
