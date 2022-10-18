package main

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zcw-admin-server/boot"
	"zcw-admin-server/global"
	"zcw-admin-server/router"
)

func main() {
	// 服务结束前关闭链接
	defer boot.Close()

	// 装载路由
	r := router.Server()

	// HTTP配置
	server := &http.Server{
		Addr:           global.CONFIG.App.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.LOG.Fatal("listen: %s\n", zap.Error(err))
		}
		global.LOG.Info("服务已开启" + global.CONFIG.App.Port)
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	global.LOG.Info("关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		global.LOG.Fatal("服务关闭原因:", zap.Error(err))
	}

	global.LOG.Info("服务退出")

}
