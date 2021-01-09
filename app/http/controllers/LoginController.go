package controllers

import (
	"net/http"

	"gitee.com/pangxianfei/frame/config"
	"gitee.com/pangxianfei/frame/helpers"
	"gitee.com/pangxianfei/frame/helpers/m"
	"gitee.com/pangxianfei/frame/http/controller"
	"gitee.com/pangxianfei/frame/request"
	"gitee.com/pangxianfei/frame/utils/crypt"
	"gitee.com/pangxianfei/frame/utils/jwt"

	"gitee.com/pangxianfei/frame/helpers/tmaic"

	"tmaic/app/http/requests"
	"tmaic/app/models"
)

type Login struct {
	controller.BaseController
}

func (l *Login) Login(c request.Context) {
	// validate and assign requestData
	var requestData requests.UserLogin
	if !l.ValidateJSON(c, &requestData, true) {
		return
	}

	user := models.User{
		Email: &requestData.Email,
	}
	if err := m.H().First(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": helpers.L(c, "auth.login.failed_not_exist")})
		return
	}

	if !crypt.BcryptCheck(*user.Password, requestData.Password) {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": helpers.L(c, "auth.login.failed_wrong_password")})
		return
	}

	// create jwt
	newJwt := jwt.NewJWT(config.GetString("auth.sign_key"))
	username := ""
	if user.Name != nil {
		username = *user.Name
	}
	if token, err := newJwt.CreateToken(string(user.ID), username); err == nil {
		c.JSON(http.StatusOK, tmaic.V{"token": token})
		return
	}

	c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": helpers.L(c, "auth.login.failed_token_generate_error")})
	return
}
