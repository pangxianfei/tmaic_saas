package groups

import (
	"gitee.com/pangxianfei/frame/route"

	"tmaic/app/http/controllers"
)

type UserAffiliationGroup struct {
	UserAffiliationController controllers.UserAffiliation
}

func (uaffg *UserAffiliationGroup) Group(group route.Grouper) {
	group.GET("/all", uaffg.UserAffiliationController.RenderAll).Name("RenderAll")
}
