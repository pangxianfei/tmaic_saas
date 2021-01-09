package listeners

import (
	"gitee.com/pangxianfei/frame/helpers/log"
	"gitee.com/pangxianfei/frame/hub"
	"github.com/golang/protobuf/proto"

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

func (user *Test) Name() hub.ListenerName {
	return "add-test"
}

func (user *Test) Subscribe() (eventPtrList []hub.Eventer) {
	log.Debug("Subscribe-test")
	return []hub.Eventer{
		&events.Test{},
	}

}

func (user *Test) Construct(paramPtr proto.Message) error {
	/** 第一执行这里
	  业务代码
	*/
	log.Debug("Construct-test")
	return nil
}
func (user *Test) Handle() error {
	/**第二执行这里
	  业务代码
	*/
	log.Info(user.user.ID)

	return nil
}
