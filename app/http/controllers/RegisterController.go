package controllers

import (
	"errors"
	"net/http"

	"gitee.com/pangxianfei/frame/helpers/log"
	"gitee.com/pangxianfei/frame/request"

	"gitee.com/pangxianfei/frame/hub"

	"gitee.com/pangxianfei/frame/config"
	"gitee.com/pangxianfei/frame/helpers"
	"gitee.com/pangxianfei/frame/helpers/m"
	"gitee.com/pangxianfei/frame/http/controller"
	"gitee.com/pangxianfei/frame/model/helper"
	"gitee.com/pangxianfei/frame/utils/crypt"
	"gitee.com/pangxianfei/frame/utils/jwt"

	"tmaic/app/events"
	pbs "tmaic/app/events/protocol_model"
	"tmaic/app/http/requests"
	"tmaic/app/models"

	"gitee.com/pangxianfei/frame/helpers/tmaic"
)

type Register struct {
	controller.BaseController
}

func (r *Register) Register(c request.Context) {

	var requestData requests.UserRegister
	if !r.ValidateJSON(c, &requestData, true) {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			responseErr, ok := err.(error)
			if ok {
				c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": responseErr.Error()})
				return
			}
			panic(err)
		}
	}()

	var token string
	var userId int64
	m.Transaction(func(TransactionHelper *helper.Helper) {
		// determine if exist
		user := models.User{
			Email: &requestData.Email,
		}
		if TransactionHelper.Exist(&user, true) {
			panic(errors.New(helpers.L(c, "auth.register.failed_existed")))
		}

		// create user
		// encrypt password //@todo move to model setter later
		encryptedPassword := crypt.Bcrypt(requestData.Password)
		user.Password = &encryptedPassword
		if err := TransactionHelper.Create(&user); err != nil {
			panic(errors.New(helpers.L(c, "auth.register.failed_system_error")))
		}

		// create jwt
		newJwt := jwt.NewJWT(config.GetString("auth.sign_key"))
		username := ""
		if user.Name != nil {
			username = *user.Name
		}
		var err error
		token, err = newJwt.CreateToken(string(user.ID), username)
		if err != nil {
			panic(helpers.L(c, "auth.register.failed_token_generate_error"))
		}

		userId = user.ID
	}, 1)

	// emit user-registered event
	ur := events.UserRegistered{}
	param := &pbs.UserRegistered{
		UserId:              uint32(userId),
		AffiliationFromCode: "",
	}
	if requestData.AffiliationFromCode == "" {
		param.AffiliationFromCode = requestData.AffiliationFromCode
	}
	ur.SetParam(param)
	if errs := hub.Emit(&ur); errs != nil {
		log.Info("user registered event emit failed", tmaic.V{"event": ur, "errors": errs})
	}

	c.JSON(http.StatusOK, tmaic.V{"token": token})
	return
}
