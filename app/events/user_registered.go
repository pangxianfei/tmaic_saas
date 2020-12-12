package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/pangxianfei/framework/hub"
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
