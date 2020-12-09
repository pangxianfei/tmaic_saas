package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/pangxianfei/framework/hub"
	pbs "tmaic/app/events/protocol_buffers"
)

func init() {
	hub.Make(&UserRegistered{})
}

type UserRegistered struct {
	hub.Event
}

func (ur *UserRegistered) ParamProto() proto.Message {
	return &pbs.UserRegistered{}
}
