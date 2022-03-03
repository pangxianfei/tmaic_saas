package middleware

import (
	"gitee.com/pangxianfei/frame/request"

	"tmaic/app/models"
	services "tmaic/app/service"
)

func TenantDbMiddle() request.HandlerFunc {
	return func(c request.Context) {
		userid, err := c.UserId()
		if err != nil {
			panic("userid is not null")
		}
		services.UserTokenService.GetCurrent(c)
		if c.ScanUserWithJSON() {
			return
		}
		users := c.User().Value().(*models.User)
		c.Set("users", users)
		c.Set("userid", userid)
		c.Set("TenantId", users.TenantsId)
		c.Next()
	}
}
