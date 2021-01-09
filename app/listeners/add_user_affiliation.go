package listeners

import (
	"errors"

	"gitee.com/pangxianfei/frame/helpers/log"

	"tmaic/app/events/protocol_model"

	"github.com/golang/protobuf/proto"

	"gitee.com/pangxianfei/frame/hub"

	"tmaic/app/events"
	"tmaic/app/models"
)

func init() {
	hub.Register(&AddUserAffiliation{})
}

type AddUserAffiliation struct {
	user                models.User
	affiliationFromCode *string
	hub.Listen
}

func (auaff *AddUserAffiliation) Name() hub.ListenerName {
	return "add-user-affiliation"
}

func (auaff *AddUserAffiliation) Subscribe() (eventPtrList []hub.Eventer) {
	return []hub.Eventer{
		&events.UserRegistered{},
	}
}

func (auaff *AddUserAffiliation) Construct(paramPtr proto.Message) error {
	// event type assertion
	_, ok := paramPtr.(*listenmodel.UserRegistered)
	if !ok {
		return errors.New("listener param is invalid")
	}
	log.Debug("test")
	return nil
}

func (userAdd *AddUserAffiliation) Handle() error {

	return nil
}
