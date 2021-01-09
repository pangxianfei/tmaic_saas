package groups

import (
	"gitee.com/pangxianfei/frame/route"

	"tmaic/app/http/controllers"
)

type AaasGroup struct {
	SassController controllers.Saas
}

func (saas *AaasGroup) Group(group route.Grouper) {

	group.POST("/index", saas.SassController.Index).Name("saas.index")

}
