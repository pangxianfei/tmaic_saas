package groups

import (
	"github.com/pangxianfei/framework/route"
	"tmaic/app/http/controllers"
)

type UserGroup struct {
	UserController controllers.User
}

func (ug *UserGroup) Group(group route.Grouper) {
	//group.GET("/info/userId", ug.UserController.Info).Can(policies.NewUserPolicy(), policy.ActionView)
	group.POST("/info", ug.UserController.Info).Name("uesrinfo")

	group.POST("/update", ug.UserController.Update).Name("user.update")
	group.GET("/delete", ug.UserController.Delete).Name("user.delete")
	group.GET("/delete-transaction", ug.UserController.DeleteTransaction).Name("user.delete-transaction")
	group.GET("/logout", ug.UserController.LogOut).Name("user.logout")
	group.GET("/restore", ug.UserController.Restore).Name("user.restore")
}
