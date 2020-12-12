package controllers

import (
	"context"
	tmaic "github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/gconv"
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/http/middleware"
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request"
	"net/http"
	"time"
	"tmaic/app/http/requests"
	"tmaic/app/models"
	"tmaic/app/policies"
	UserService "tmaic/app/service/users"
)

type User struct {
	controller.BaseController
}

func (*User) LogOut(c request.Context) {
	if err := middleware.Revoke(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tmaic.V{})
	return
}

func (u *User) Info(c request.Context) {
	if c.ScanUserWithJSON() {
		return
	}

	user := c.User().Value().(*models.User)

	if permit, _ := u.Authorize(c, policies.NewUserPolicy(), policy.ActionView); !permit {
		c.JSON(http.StatusForbidden, tmaic.V{"error": policy.UserNotPermitError{}.Error()})
		return
	}

	c.JSON(http.StatusOK, tmaic.V{"data": user})
	return
}

func (*User) AllUser(c request.Context) {
	user := &models.User{}

	c.JSON(http.StatusOK, tmaic.V{"data": user})
	return
}

func (*User) PaginateUser(c request.Context) {
	user := &models.User{}

	c.JSON(http.StatusOK, tmaic.V{"data": user})
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

	//更新
	t := UserService.UserService.Get(gconv.Int64(requestData.ID))
	t.Email = requestData.Email
	t.Name = requestData.Name
	_ = UserService.UserService.Update(t)

	user := &models.User{
		ID: 19,
	}

	//查询
	if err := user.DB().Debug().Where(user).First(user).Error; err != nil {
		c.JSON(http.StatusOK, tmaic.Output{"data": err})
		return
	}

	//查询
	DeleteUser := models.User{
		ID: 30,
	}
	if err := user.DB().Debug().Delete(DeleteUser).Error; err != nil {
		c.JSON(http.StatusOK, tmaic.Output{"data": err})
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	ctxDB := user.DB().WithContext(ctx)

	for i := 0; i < 100; i++ {
		go ctxDB.Debug().First(&user)
	}

	c.JSON(http.StatusOK, tmaic.Output{"data": t})
	return
}

func (*User) Delete(c request.Context) {

	user := models.User{
		ID: 78,
	}
	if err := m.H().Delete(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err})
		return
	}
	c.JSON(http.StatusOK, tmaic.V{"data": true})
	return
}
func (*User) DeleteTransaction(c request.Context) {
	defer func() { // handle transaction error
		if err := recover(); err != nil {
			c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err.(error).Error()})
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

	c.JSON(http.StatusOK, tmaic.V{"data": true})
	return
}
func (*User) Restore(c request.Context) {

	modifyUser := models.User{
		ID: 78,
	}

	if err := m.H().Restore(&modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err})
		return
	}
	c.JSON(http.StatusOK, tmaic.V{"data": true})
	return
}
