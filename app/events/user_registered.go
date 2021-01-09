package events

import (
	"github.com/golang/protobuf/proto"

	"gitee.com/pangxianfei/frame/hub"
	pbs "tmaic/app/events/events_buffers"
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
