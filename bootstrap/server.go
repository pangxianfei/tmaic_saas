package bootstrap

import (
	"context"
	c "gitee.com/pangxianfei/frame/config"
	"gitee.com/pangxianfei/frame/helpers/log"
	"gitee.com/pangxianfei/frame/helpers/tmaic"
	"gitee.com/pangxianfei/frame/helpers/zone"
	"gitee.com/pangxianfei/frame/request"
	"gitee.com/pangxianfei/frame/sentry"
	"net/http"
	"sync"
	"tmaic/resources/views"
	"tmaic/routes"
)

func HttpServe(parentCtx context.Context, wg *sync.WaitGroup) {
	r := request.New()

	sentry.Use(r.GinEngine(), false)

	Middleware(r)

	routes.Register(r)

	views.Initialize(r)

	s := &http.Server{
		Addr:           ":" + c.GetString("app.port"),
		Handler:        r,
		ReadTimeout:    zone.Duration(c.GetInt64("app.read_timeout_seconds")) * zone.Second,
		WriteTimeout:   zone.Duration(c.GetInt64("app.write_timeout_seconds")) * zone.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		log.Info("Served At", tmaic.V{"Addr": s.Addr})
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()
	<-parentCtx.Done()
	//log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(parentCtx, 5*zone.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", tmaic.V{"error": err})
	}

	wg.Done()
}
