package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	c "github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework/graceful"
	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/helpers/toto"
	"github.com/pangxianfei/framework/helpers/zone"
	"github.com/pangxianfei/framework/monitor"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/sentry"

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
		log.Info("system call", toto.V{"call": call})
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
		log.Info("Served At", toto.V{"Addr": s.Addr})
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	<-parentCtx.Done()

	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(parentCtx, 5*zone.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", toto.V{"error": err})
	}

	wg.Done()
}
