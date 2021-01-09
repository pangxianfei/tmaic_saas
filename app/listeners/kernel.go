package listeners

import "gitee.com/pangxianfei/frame/hub"

func Initialize() {
	// 初始化主题和频道
	hub.RegisterQueue()
}
