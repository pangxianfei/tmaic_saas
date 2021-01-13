package groups

import (
	"gitee.com/pangxianfei/frame/route"

	"tmaic/app/http/controllers"
)

type Events struct {
	EventsController controllers.Events
}

// Group bt路由组
func (ev *Events) Group(group route.Grouper) {

	// 磁盘使用情况
	group.POST("/demo", ev.EventsController.EventsDemo).Name("EventsDemo")

}
