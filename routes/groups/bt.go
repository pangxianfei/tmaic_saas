package groups

import (
	"gitee.com/pangxianfei/frame/route"

	"tmaic/app/http/controllers"
)

type BtGroup struct {
	BtController controllers.Bt
}

// Group bt路由组
func (bt *BtGroup) Group(group route.Grouper) {

	// 磁盘使用情况
	group.POST("/GetDiskInfo", bt.BtController.GetDiskInfo).Name("GetDiskInfo")

}
