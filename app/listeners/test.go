package listeners

import (
	"gitee.com/pangxianfei/frame/helpers/log"
	"gitee.com/pangxianfei/frame/hub"
	"github.com/golang/protobuf/proto"

	"tmaic/app/events"
	listenmodel "tmaic/app/events/protocol_model"
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
	  业务代码 这里通常为 Handle 数据作为准备的工作
	*/
	test, _ := paramPtr.(*listenmodel.Test)
	log.Debug("Construct-userid:")
	log.Debug(test.Id)
	// 出始化 user 数据 handle 直接拿user 数据
	user.user.ID = int64(test.Id)

	return nil
}
func (user *Test) Handle() error {
	/**第二执行这里
	  业务代码
	*/
	log.Info(user.user.ID)

	return nil
}
