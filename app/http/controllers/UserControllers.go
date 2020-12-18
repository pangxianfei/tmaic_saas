package controllers

import (
	. "github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/http/middleware"
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request"
	"github.com/spf13/cast"
	"net/http"

	"tmaic/app/http/requests"
	"tmaic/app/models"
	"tmaic/app/policies"
	"tmaic/app/service/users"
)

type User struct {
	controller.BaseController
}

func (*User) LogOut(c request.Context) {
	if err := middleware.Revoke(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, V{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, V{})
	return
}

func (u *User) Info(c request.Context) {
	if c.ScanUserWithJSON() {
		return
	}

	user := c.User().Value().(*models.User)

	if permit, _ := u.Authorize(c, policies.NewUserPolicy(), policy.ActionView); !permit {
		c.JSON(http.StatusForbidden, V{"error": policy.UserNotPermitError{}.Error()})
		return
	}

	c.JSON(http.StatusOK, V{"data": user})
	return
}

func (*User) AllUser(c request.Context) {
	user := &models.User{}

	c.JSON(http.StatusOK, V{"data": user})
	return
}

func (*User) PaginateUser(c request.Context) {
	user := &models.User{}

	c.JSON(http.StatusOK, V{"data": user})
	return
}

type APIUser struct {
	ID   uint
	Name string
}

func (u *User) Update(c request.Context) {
	var requestData requests.UserUpdateRequests

	if !u.ValidateJSON(c, &requestData, true) {
		return
	}

	// 更新
	t := users.UserService.Get(cast.ToInt64(requestData.ID))
	t.Email = requestData.Email
	t.Name = requestData.Name
	_ = users.UserService.Update(t)

	var id int64
	id = 18
	userUpdate := models.User{
		ID: id,
	}
	if err := m.H().First(&userUpdate, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, V{"error": err.Error()})
		return
	}

	modifyUser := models.User{
		Name: requestData.Name,
	}

	if err := m.H().Save(&userUpdate, modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, V{"error": err})
		return
	}

	// 查询
	/*
		user := &models.User{
			ID: 18,
		}
		if err := user.DB().Debug().Where(user).First(user).Error; err != nil {
			c.JSON(http.StatusOK, Output{"data": err})
			return
		}*/

	// 删除
	/*
		DeleteUser := models.User{
			ID: 30,
		}
		if err := user.DB().Debug().Delete(DeleteUser).Error; err != nil {
			c.JSON(http.StatusOK, Output{"data": err})
			return
		}*/

	// 批量查询
	/*
		ctx, _ := context.WithTimeout(context.Background(), time.Second)

		ctxDB := user.DB().WithContext(ctx)

		for i := 0; i < 100; i++ {
			go ctxDB.Debug().First(&user)
		}
	*/

	c.JSON(http.StatusOK, Output{"data": userUpdate})
	return
}

func (*User) Delete(c request.Context) {

	user := models.User{
		ID: 78,
	}
	if err := m.H().Delete(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, V{"error": err})
		return
	}
	c.JSON(http.StatusOK, V{"data": true})
	return
}
func (*User) DeleteTransaction(c request.Context) {
	defer func() { // handle transaction error
		if err := recover(); err != nil {
			c.JSON(http.StatusUnprocessableEntity, V{"error": err.(error).Error()})
			return
		}
	}()

	user := models.User{
		ID: 78,
	}
	m.Transaction(func(h *m.Helper) {
		user.SetTX(h.DB()) // important
		if err := h.Delete(&user, false); err != nil {
			panic(err)
		}
	}, 1)

	c.JSON(http.StatusOK, V{"data": true})
	return
}
func (*User) Restore(c request.Context) {

	modifyUser := models.User{
		ID: 78,
	}

	if err := m.H().Restore(&modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, V{"error": err})
		return
	}
	c.JSON(http.StatusOK, V{"data": true})
	return
}
