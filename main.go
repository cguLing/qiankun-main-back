package main

import (
	"bus-backend-go/autoinit"
	"bus-backend-go/conf"
	ginAutoRoute "bus-backend-go/route"
	"bus-backend-go/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// TODO 使用es时放开 新的初始化方式
	autoinit.InitTouch()
	router := gin.New()
	router.Use(gin.Recovery())
	ginAutoRoute.RegisterRoute(router, utils.Log)
	srv := &http.Server{
		Addr:    ":" + conf.Sysconfig.Port,
		Handler: router,
	}

	// 启动服务器
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.Log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utils.Log.Error("Server Shutdown: ", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		utils.Log.Info("timeout of 5 seconds.")
	}
	utils.Log.Info("Server exiting")
}