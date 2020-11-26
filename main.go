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

// @caution cannot use config methods to get config in init function
func main() {
	//j := &jobs.ExampleJob{}
	//j.SetParam(&pbs.ExampleJob{Query: "test", PageNumber: 111, ResultPerPage: 222})
	////j.SetDelay(5 * zone.Second)
	//err := job.Dispatch(j)
	//fmt.Println(err)

	//go hub.On("add-user-affiliation")  // go run artisan.go queue:listen add-user-affiliation

	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
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

	// tmaic framework shutdown
	graceful.ShutDown(false)

	log.Info("Server exited")
}

func httpServe(parentCtx context.Context, wg *sync.WaitGroup) {
	r := request.New()

	sentry.Use(r.GinEngine(), false)

	bootstrap.Middleware(r)

	//r.Use(middleware.IUser(&models.YourUserModel{})) // set default auth user model, or use config auth.model_ptr

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

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	ctx, cancel := context.WithTimeout(parentCtx, 5*zone.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", toto.V{"error": err})
	}

	wg.Done()
}
