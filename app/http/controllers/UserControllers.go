package controllers

import (
	"net/http"

	tmaic "github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/http/middleware"
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request"

	"tmaic/app/models"
	"tmaic/app/policies"
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

func (*User) Update(c request.Context) {
	var id uint
	id = 78
	user := models.User{
		ID: &id,
	}
	if err := m.H().First(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err.Error()})
		return
	}

	name := "pangxianfei"
	modifyUser := models.User{
		Name: &name,
	}
	if err := m.H().Save(&user, modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err})
		return
	}
	c.JSON(http.StatusOK, tmaic.V{"data": user})
	return

}
func (*User) Delete(c request.Context) {
	var id uint
	id = 78
	user := models.User{
		ID: &id,
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

	var id uint
	id = 78
	user := models.User{
		ID: &id,
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
	var id uint
	id = 78
	modifyUser := models.User{
		ID: &id,
	}

	if err := m.H().Restore(&modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err})
		return
	}
	c.JSON(http.StatusOK, tmaic.V{"data": true})
	return
}
