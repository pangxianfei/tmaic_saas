package groups

import (
	"gitee.com/pangxianfei/frame/route"

	"tmaic/app/http/controllers"
)

type BtGroup struct {
	BtController controllers.Bt
}

func (bt *BtGroup) Group(group route.Grouper) {

	// 磁盘使用情况
	group.POST("/bt/GetDiskInfo", bt.BtController.GetDiskInfo)

}
