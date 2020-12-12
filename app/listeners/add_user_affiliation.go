package listeners

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/hub"
	"tmaic/app/events"
	pbs "tmaic/app/events/protocol_buffers"
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

func (userAdd *AddUserAffiliation) Name() hub.ListenerName {
	return "add-user-affiliation"
}

func (userAdd *AddUserAffiliation) Subscribe() (eventPtrList []hub.Eventer) {

	return []hub.Eventer{
		&events.UserRegistered{},
	}

}

func (userAdd *AddUserAffiliation) Construct(paramPtr proto.Message) error {
	// 事件类型断言
	param, ok := paramPtr.(*pbs.UserRegistered)
	if !ok {
		return errors.New("listener param is invalid")
	}

	userAdd.affiliationFromCode = nil
	if param.AffiliationFromCode != "" && checkFromCode(param.AffiliationFromCode) {
		userAdd.affiliationFromCode = &param.AffiliationFromCode
	}

	uid := int64(param.GetUserId())
	userAdd.user = models.User{ID: uid}
	if err := m.H().First(&userAdd.user, false); err != nil {
		return err
	}

	return nil
}

func (userAdd *AddUserAffiliation) Handle() error {
	// 添加用户从属关系
	if config.GetBool("user_affiliation.enable") {
		uaffPtr := &models.UserAffiliation{
			UserID: &userAdd.user.ID,
		}
		var err error
		if userAdd.affiliationFromCode != nil {
			err = uaffPtr.InsertNode(&userAdd.user, *userAdd.affiliationFromCode)
		} else {
			err = uaffPtr.InsertNode(&userAdd.user)
		}
		if err != nil {
			return errors.New("user affiliation insert failed")
		}
	}

	return nil
}

// 检验验证码
func checkFromCode(affiliationFromCode string) bool {
	FromCode := models.UserAffiliation{
		Code: &affiliationFromCode,
	}
	return m.H().Exist(&FromCode, false)
}
