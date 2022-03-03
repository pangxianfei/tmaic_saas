package services

import (
	// "gitee.com/pangxianfei/frame/helpers/debug"
	"gitee.com/pangxianfei/frame/model/sysmodel"
	"gitee.com/pangxianfei/frame/request"

	"tmaic/app/models"
	"tmaic/app/repositories"
)

var UserTokenService = newUserTokenService()

func newUserTokenService() *userTokenService {
	return &userTokenService{}
}

type userTokenService struct {
}

func (s *userTokenService) Get(id int64) *sysmdel.UserToken {
	var db sysmdel.UserToken
	return repositories.UserTokenRepository.Get(db.DB(), id)
}

func (s *userTokenService) Take(where ...interface{}) *sysmdel.UserToken {
	var db sysmdel.UserToken

	return repositories.UserTokenRepository.Take(db.DB(), where...)

}

// GetCurrentUserId 获取当前登录用户的id
func (s *userTokenService) GetCurrentUserId(ctx request.Context) int64 {
	userid, err := ctx.UserId()
	if err == nil {
		return int64(userid)
	}
	return 0
}

// GetCurrent 获取当前登录用户
func (s *userTokenService) GetCurrent(c request.Context) *models.User {
	token, _ := c.Get("CONTEXT_TOKEN_KEY")

	if token == "" {
		panic("token is not null")
	}

	userid, err := c.UserId()
	if err != nil {
		panic("userid is not null")
	}

	var u models.User
	user := &models.User{
		ID: int64(userid),
	}

	u.DB().Take(&user)

	var T sysmdel.Tenants

	TenantsModel := sysmdel.Tenants{
		TenantsId: user.TenantsId,
	}

	T.DB().Where(&TenantsModel).Find(&TenantsModel)

	c.Set("TenantsId", user.TenantsId)
	c.Set("TenantsDB", TenantsModel)

	return user
}
