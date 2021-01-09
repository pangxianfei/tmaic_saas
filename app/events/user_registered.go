package events

import (
	"gitee.com/pangxianfei/frame/hub"
	"github.com/golang/protobuf/proto"

	event "tmaic/app/events/protocol_model"
)

func init() {
	hub.Make(&UserRegistered{})
}

type UserRegistered struct {
	hub.Event
	UserId uint32
}

func (ur *UserRegistered) ParamProto() proto.Message {
	return &event.UserRegistered{}
}
