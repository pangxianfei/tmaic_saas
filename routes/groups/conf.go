package groups

import (
	"github.com/pangxianfei/framework/route"
	"tmaic/app/http/controllers"
)

type ConfGroup struct {
	ConfController controllers.Conf
}

func (ug *ConfGroup) Group(group route.Grouper) {
	group.GET("/conf", ug.ConfController.Index).Name("conf.index")
}
