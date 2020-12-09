package middleware

import (
	tmaic "github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/helpers/zone"
	"github.com/pangxianfei/framework/request"
)

func Example() request.HandlerFunc {
	return func(c request.Context) {
		t := zone.Now()
		c.Set("test", "test")
		c.Next()
		// after request
		latency := zone.Since(t)
		log.Info("latency", tmaic.V{"latency": latency})

	}
}
