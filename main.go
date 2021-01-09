package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	c "gitee.com/pangxianfei/frame/config"
	"gitee.com/pangxianfei/frame/graceful"
	"gitee.com/pangxianfei/frame/helpers/log"
	"gitee.com/pangxianfei/frame/helpers/zone"
	"gitee.com/pangxianfei/frame/monitor"
	"gitee.com/pangxianfei/frame/request"
	"gitee.com/pangxianfei/frame/sentry"

	"gitee.com/pangxianfei/frame/helpers/tmaic"

	"tmaic/bootstrap"
	"tmaic/resources/views"
	"tmaic/routes"
)

func init() {
	bootstrap.Initialize()
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		call := <-quit
		log.Info("system call", tmaic.V{"call": call})
		cancel()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go httpServe(ctx, wg)
	wg.Add(1)
	go monitor.HttpMonitorServe(ctx, wg)

	wg.Wait()

	graceful.ShutDown(false)

	log.Info("Server exited")
}

func httpServe(parentCtx context.Context, wg *sync.WaitGroup) {
	r := request.New()

	sentry.Use(r.GinEngine(), false)

	bootstrap.Middleware(r)

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

	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(parentCtx, 5*zone.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", tmaic.V{"error": err})
	}

	wg.Done()
}
