package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/pangxianfei/framework/hub"
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
