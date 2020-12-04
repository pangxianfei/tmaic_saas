package listeners

import "github.com/pangxianfei/framework/hub"

func Initialize() {
	// 初始化主题和频道
	hub.RegisterQueue()
}
