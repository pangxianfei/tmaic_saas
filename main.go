package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gitee.com/pangxianfei/frame/graceful"
	"gitee.com/pangxianfei/frame/helpers/log"
	"gitee.com/pangxianfei/frame/helpers/tmaic"
	"gitee.com/pangxianfei/frame/monitor"

	"tmaic/bootstrap"
)

func init() {
	bootstrap.Initialize()
}

func main() {
	// 父context(利用根context得到)
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
	go bootstrap.HttpServe(ctx, wg)

	wg.Add(1)
	go monitor.HttpMonitorServe(ctx, wg)

	wg.Wait()

	graceful.ShutDown(false)

	log.Info("Server exited")
}
