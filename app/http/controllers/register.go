package controllers

import (
	"errors"
	"net/http"

	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/helpers/toto"
	"github.com/pangxianfei/framework/request"

	"github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework/helpers"
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/model/helper"
	"github.com/pangxianfei/framework/utils/crypt"
	"github.com/pangxianfei/framework/utils/jwt"

	"tmaic/app/http/requests"
	"tmaic/app/models"
)

type Register struct {
	controller.BaseController
}

func (r *Register) Register(c request.Context) {
	// validate and assign requestData
	var requestData requests.UserRegister
	if !r.ValidateJSON(c, &requestData, true) {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			responseErr, ok := err.(error)
			if ok {
				c.JSON(http.StatusUnprocessableEntity, toto.V{"error": responseErr.Error()})
				return
			}
			panic(err)
		}
	}()

	var token string
	var userId uint
	m.Transaction(func(TransactionHelper *helper.Helper) {

		user := models.User{
			Name: &requestData.Email,
			Email: &requestData.Email,
		}
		if TransactionHelper.Exist(&user, true) {
			panic(errors.New(helpers.L(c, "auth.register.failed_existed")))
		}


		encryptedPassword := crypt.Bcrypt(requestData.Password)
		user.Password = &encryptedPassword
		if err := TransactionHelper.Create(&user); err != nil {
			panic(errors.New(helpers.L(c, "auth.register.failed_system_error")))
		}

		// create token
		newJwt := jwt.NewJWT(config.GetString("auth.sign_key"))
		username := ""
		if user.Name != nil {
			username = *user.Name
		}
		var err error
		token, err = newJwt.CreateToken(string(*user.ID), username)
		if err != nil {
			panic(helpers.L(c, "auth.register.failed_token_generate_error"))
		}

		userId = *user.ID
	}, 1)


	log.Info("user registered event emit failed", toto.V{"token": token})
	c.JSON(http.StatusOK, toto.V{"token": token})
	return
}
