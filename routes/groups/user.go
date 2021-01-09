package groups

import (
	"gitee.com/pangxianfei/frame/policy"
	"gitee.com/pangxianfei/frame/route"

	"tmaic/app/http/controllers"
	"tmaic/app/policies"
)

type UserGroup struct {
	UserController controllers.User
}

func (ug *UserGroup) Group(group route.Grouper) {
	group.GET("/info/:userId", ug.UserController.Info).Name("user.info").Can(policies.NewUserPolicy(), policy.ActionView)

	group.GET("/update", ug.UserController.Update).Name("user.update")
	group.GET("/delete", ug.UserController.Delete).Name("user.delete")
	group.GET("/delete-transaction", ug.UserController.DeleteTransaction).Name("DeleteTransaction")
	group.GET("/logout", ug.UserController.LogOut).Name("logout")
	group.GET("/restore", ug.UserController.Restore).Name("restore")
}
