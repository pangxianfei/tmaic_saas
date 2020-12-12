package listeners

import (
	"errors"
	"tmaic/app/events/protocol_model"

	"github.com/golang/protobuf/proto"
	"github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/hub"
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
	param, ok := paramPtr.(*listenmodel.UserRegistered)
	if !ok {
		return errors.New("listener param is invalid")
	}

	auaff.affiliationFromCode = nil
	if param.AffiliationFromCode != "" && checkFromCode(param.AffiliationFromCode) {
		auaff.affiliationFromCode = &param.AffiliationFromCode
	}

	uid := int64(param.GetUserId())
	auaff.user = models.User{ID: uid}
	if err := m.H().First(&auaff.user, false); err != nil {
		return err
	}

	return nil
}

func (userAdd *AddUserAffiliation) Handle() error {
	// add user affiliation
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
