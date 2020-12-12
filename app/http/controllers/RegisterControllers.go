package controllers

import (
	"errors"
	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/hub"
	"net/http"
	"strconv"
	"tmaic/app/events"
	pbs "tmaic/app/events/protocol_model"
	"tmaic/app/http/requests"
	"tmaic/app/models"

	tmaic "github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/helpers"
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/model/helper"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/utils/crypt"
	"github.com/pangxianfei/framework/utils/jwt"
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
				c.JSON(http.StatusUnprocessableEntity, tmaic.V{"error": responseErr.Error()})
				return
			}
			panic(err)
		}
	}()

	var token string
	var userId uint
	m.Transaction(func(TransactionHelper *helper.Helper) {

		user := models.User{
			Name:  requestData.Email,
			Email: requestData.Email,
		}
		if TransactionHelper.Exist(&user, true) {
			panic(errors.New(helpers.L(c, "auth.register.failed_existed")))
		}

		encryptedPassword := crypt.Bcrypt(requestData.Password)
		user.Password = encryptedPassword
		if err := TransactionHelper.Create(&user); err != nil {
			panic(errors.New(helpers.L(c, "auth.register.failed_system_error")))
		}

		// create token
		newJwt := jwt.NewJWT()
		username := ""
		if user.Name != "" {
			username = user.Name
		}
		var err error
		token, err = newJwt.CreateToken(strconv.Itoa(int(user.ID)), username)
		if err != nil {
			panic(helpers.L(c, "auth.register.failed_token_generate_error"))
		}

		userId = uint(user.ID)
	}, 1)

	// 注册事件
	ur := events.UserRegistered{}
	param := &pbs.UserRegistered{
		UserId:              uint32(userId),
		AffiliationFromCode: "",
	}
	if requestData.AffiliationFromCode != nil {
		//验证码
		param.AffiliationFromCode = *requestData.AffiliationFromCode
	}
	ur.SetParam(param)
	if errs := hub.Emit(&ur); errs != nil {
		log.Info("user registered event emit failed", tmaic.Output{"event": ur, "errors": errs})
	}

	c.JSON(http.StatusOK, tmaic.Output{"token": token})
	return
}
