package controllers

import (
	"github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/helpers"
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/utils/crypt"
	"github.com/pangxianfei/framework/utils/jwt"
	"net/http"
	"tmaic/app/http/requests"
	"tmaic/app/models"
)

type Login struct {
	controller.BaseController
}

func (l *Login) Login(c request.Context) {

	var requestData requests.UserLogin
	if !l.ValidateJSON(c, &requestData, true) {
		return
	}

	user := models.User{
		Email: requestData.Email,
	}
	if err := m.H().First(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": helpers.L(c, "auth.login.failed_not_exist")})
		return
	}

	if !crypt.BcryptCheck(user.Password, requestData.Password) {
		c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": helpers.L(c, "auth.login.failed_wrong_password")})
		return
	}

	// create jwt
	newJwt := jwt.NewJWT()

	if token, err := newJwt.CreateToken(string(user.ID), user.Name); err == nil {
		c.JSON(http.StatusOK, tmaic.Output{"token": token})
		return
	}

	c.JSON(http.StatusUnprocessableEntity, tmaic.Output{"error": helpers.L(c, "auth.login.failed_token_generate_error")})

	return
}
