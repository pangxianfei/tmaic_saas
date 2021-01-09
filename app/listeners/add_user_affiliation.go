package listeners

import (
	"errors"

	"github.com/golang/protobuf/proto"

	"gitee.com/pangxianfei/frame/config"
	"gitee.com/pangxianfei/frame/helpers/m"
	"gitee.com/pangxianfei/frame/hub"

	"tmaic/app/events"
	pbs "tmaic/app/events/events_buffers"
	"tmaic/app/models"
)

func init() {
	hub.Register(&AddUserAffiliation{})
}

type AddUserAffiliation struct {
	user                models.User
	affiliationFromCode string
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
	param, ok := paramPtr.(*pbs.UserRegistered)

	if !ok {
		return errors.New("listener param is invalid")
	}

	if param.AffiliationFromCode != "" && checkFromCode(param.AffiliationFromCode) {
		auaff.affiliationFromCode = param.AffiliationFromCode
	}

	uid := int64(param.GetUserId())
	auaff.user = models.User{ID: uid}
	if err := m.H().First(&auaff.user, false); err != nil {
		return err
	}

	return nil
}

// Handle 方法
func (auaff *AddUserAffiliation) Handle() error {
	// add user affiliation
	if config.GetBool("user_affiliation.enable") {
		uaffPtr := &models.UserAffiliation{
			UserID: auaff.user.ID,
		}
		var err error
		if auaff.affiliationFromCode != "" {
			err = uaffPtr.InsertNode(&auaff.user, auaff.affiliationFromCode)
		} else {
			err = uaffPtr.InsertNode(&auaff.user)
		}
		if err != nil {
			return errors.New("user affiliation insert failed")
		}
	}

	return nil
}

// check affiliationFromCode is valid
func checkFromCode(affiliationFromCode string) bool {
	uaff := models.UserAffiliation{
		Code: &affiliationFromCode,
	}
	return m.H().Exist(&uaff, false)
}
