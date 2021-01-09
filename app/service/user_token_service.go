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

	// userToken := cache.UserTokenCache.Get(token)
	// 没找到授权
	/*
		if userToken == nil || userToken.Status == constants.StatusDeleted {
			return nil
		}

		// 授权过期
		if userToken.ExpiredAt <= date.NowTimestamp() {
			return nil
		}
		user := cache.UserCache.Get(userToken.UserId)

		if user == nil || user.Status != constants.StatusOk {
			return nil
		}*/

}

/*
// CheckLogin 检查登录状态
func (s *userTokenService) CheckLogin(ctx request.Context) (*model.User, *simple.CodeError) {
	user := s.GetCurrent(ctx)

}
*/
/*
// 退出登录
func (s *userTokenService) Signout(ctx request.Context) error {

}

// 从请求体中获取UserToken
func (s *userTokenService) GetUserToken(ctx iris.Context) string {

}
*/
