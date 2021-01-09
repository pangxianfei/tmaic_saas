package controllers

import (
	"net/http"

	"gitee.com/pangxianfei/frame/config"
	"gitee.com/pangxianfei/frame/helpers/m"
	"gitee.com/pangxianfei/frame/http/controller"
	"gitee.com/pangxianfei/frame/http/middleware"
	"gitee.com/pangxianfei/frame/model"
	"gitee.com/pangxianfei/frame/policy"
	"gitee.com/pangxianfei/frame/request"

	"tmaic/app/models"
	"tmaic/app/policies"

	"gitee.com/pangxianfei/frame/helpers/tmaic"
)

type User struct {
	controller.BaseController
}

func (*User) LogOut(c request.Context) {
	if err := middleware.Revoke(c, config.GetString("auth.sign_key")); err != nil {
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
	outArr, err := user.ObjArr([]model.Filter{}, []model.Sort{}, 0, false)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tmaic.V{"data": outArr.([]models.User)})
	return
}

func (*User) PaginateUser(c request.Context) {
	user := &models.User{}
	pagination, err := user.ObjArrPaginate(c, 25, []model.Filter{}, []model.Sort{}, 0, false)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tmaic.V{"data": tmaic.V{"item": pagination.ItemArr(), "totalPage": pagination.LastPage(), "currentPage": pagination.CurrentPage(), "count": pagination.Count(), "total": pagination.Total()}})
	return
}

func (u *User) Update(c request.Context) {
	var id int64
	id = 1
	user := models.User{
		ID: id,
	}
	if err := m.H().First(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err.Error()})
		return
	}

	/*
		name := "t2222es123t"
		modifyUser := models.User{
			Name: &name,
		}
		if err := m.H().Save(&user, modifyUser); err != nil {
			c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err})
			return
		}
	*/
	c.JSON(http.StatusOK, tmaic.V{"data": user})
	return

}
func (*User) Delete(c request.Context) {
	var id int64
	id = 1
	user := models.User{
		ID: id,
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

	var id int64
	id = 1
	user := models.User{
		ID: id,
	}
	m.Transaction(func(h *m.Helper) {
		user.SetTX(h.DB())
		if err := h.Delete(&user, false); err != nil {
			panic(err)
		}
	}, 1)

	c.JSON(http.StatusOK, tmaic.V{"data": true})
	return
}
func (*User) Restore(c request.Context) {
	var id int64
	id = 1
	modifyUser := models.User{
		ID: id,
	}

	if err := m.H().Restore(&modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": err})
		return
	}
	c.JSON(http.StatusOK, tmaic.V{"data": true})
	return
}
