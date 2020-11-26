package listeners

import "github.com/pangxianfei/framework/hub"

func Initialize() {
	// initialize topic and channel
	hub.RegisterQueue()
}
