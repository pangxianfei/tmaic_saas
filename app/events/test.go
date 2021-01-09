package events

import (
	"gitee.com/pangxianfei/frame/hub"
	"github.com/golang/protobuf/proto"

	events "tmaic/app/events/events_buffers"
)

func init() {
	hub.Make(&Test{})
}

type Test struct {
	hub.Event
	Id uint32
}

func (ur *Test) ParamProto() proto.Message {
	return &events.Test{}
}
