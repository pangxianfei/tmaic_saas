package controllers

import (
	"fmt"
	"time"

	"gitee.com/pangxianfei/frame/helpers/log"
	"gitee.com/pangxianfei/frame/helpers/tmaic"
	"gitee.com/pangxianfei/frame/http/controller"
	"gitee.com/pangxianfei/frame/hub"
	"gitee.com/pangxianfei/frame/request"
	"gitee.com/pangxianfei/frame/utils/md5"
	"github.com/qifengzhang007/goCurl"

	"tmaic/app/events"
	pbs "tmaic/app/events/protocol_model"
)

type Bt struct {
	controller.BaseController
}

func (Bt *Bt) GetDiskInfo(c request.Context) {

	ur := events.Test{}
	param := &pbs.Test{
		Id: 1,
	}

	ur.SetParam(param)
	if errs := hub.Emit(&ur); errs != nil {
		log.Info("user registered event emit failed", tmaic.V{"event": ur, "errors": errs})
	}

	cli := goCurl.CreateHttpClient()

	var token string
	var md5Sign string
	timestamp := time.Now().Unix()

	md5Sign = md5.CreateMd5("dK5jkigIQ03zFjxh7ilQWgQhp4OCKqQI")
	str := fmt.Sprintf("%d%s", timestamp, md5Sign)
	token = md5.CreateMd5(str)

	resp, _ := cli.Post("http://192.168.3.107:8888/system?action=GetDiskInfo", goCurl.Options{
		FormParams: map[string]interface{}{
			"request_time":  timestamp,
			"request_token": token,
		},
	})

	txt, _ := resp.GetContents()

	tmaic.Success(c, "yes", txt)
}
