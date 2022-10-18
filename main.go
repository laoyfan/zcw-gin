package main

import (
	"context"
	"fmt"
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
	defer boot.Close()
	server := router.Server()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.LOG.Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	fmt.Println(server)

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
