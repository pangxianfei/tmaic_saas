package events

import (
	"gitee.com/pangxianfei/frame/hub"
	"github.com/golang/protobuf/proto"

	event "tmaic/app/events/protocol_model"
)

func init() {
	hub.Make(&Test{})
}

type Test struct {
	hub.Event
	Id uint32
}

func (ur *Test) ParamProto() proto.Message {
	return &event.Test{}
}
