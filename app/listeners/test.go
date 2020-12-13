package listeners

import (
	"github.com/golang/protobuf/proto"
	"github.com/pangxianfei/framework/helpers/log"

	"github.com/pangxianfei/framework/hub"

	"tmaic/app/events"
	"tmaic/app/models"
)

func init() {
	hub.Register(&Test{})
}

type Test struct {
	user models.User
	hub.Listen
}

func (auaff *Test) Name() hub.ListenerName {
	return "add-test"
}

func (auaff *Test) Subscribe() (eventPtrList []hub.Eventer) {
	log.Debug("Subscribe-test")
	return []hub.Eventer{
		&events.Test{},
	}

}

func (auaff *Test) Construct(paramPtr proto.Message) error {
	/** 第一执行这里
	  业务代码
	*/
	log.Debug("Construct-test")
	return nil
}
func (userAdd *Test) Handle() error {
	/**第二执行这里
	  业务代码
	*/
	log.Debug("Handle-test")
	return nil
}
