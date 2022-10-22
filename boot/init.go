package boot

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zcw-admin-server/boot/internal"
	"zcw-admin-server/global"
	"zcw-admin-server/middleware"
	"zcw-admin-server/util"
)

const ConfigDir = "config" //配置文件夹

// 开启服务前初始化准备

func init() {
	// 读取ConfigDir目录下所有配置
	fileNames := util.GetPathFileNames(ConfigDir) // 获取config文件夹下配置文件名称
	if len(fileNames) > 0 {
		for _, fileName := range fileNames {
			v := viper.New()                            // 每个文件定义一个实例
			v.SetConfigFile(ConfigDir + "/" + fileName) // 配置文件
			if err := v.ReadInConfig(); err != nil {    // 读取文件 此时异常需要panic
				panic(fmt.Errorf("读取配置文件异常: %s \n", err))
			}
			v.WatchConfig()
			v.OnConfigChange(func(e fsnotify.Event) {
				fmt.Println("配置文件变更:", e.Name)
				if err := v.Unmarshal(&global.CONFIG); err != nil { // 配置写入
					panic(fmt.Errorf("Config转换异常: %s \n", err))
				}
				created()
			})
			if err := v.Unmarshal(&global.CONFIG); err != nil { // 配置写入
				panic(fmt.Errorf("Config转换异常: %s \n", err))
			}
		}
	}
	created()
}

// 初始化连接

func created() {
	internal.Zap()
	internal.Mysql()
	internal.Redis()
	internal.Limiter()
	internal.Valid()
}

// 关闭开启的连接资源

func Close() {
	internal.MysqlClose()
	internal.RedisClose()
}

// 开启服务

func Server() {
	// 服务结束前关闭链接
	defer Close()
	// 关闭控制台颜色
	gin.DisableConsoleColor()
	// 设置模式
	gin.SetMode(global.CONFIG.App.Mode)
	// 开启gin实例
	r := gin.New()
	// 全局处理中间件
	r.Use(
		middleware.Logger,  //日志处理
		middleware.Cors,    //跨域处理
		middleware.Error,   //异常处理
		middleware.Limiter, //限流处理
	)
	// HTTP配置
	server := &http.Server{
		Addr:           global.CONFIG.App.Port,
		Handler:        internal.Route(r),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 开启服务
	go func() {
		global.LOG.Info("服务开启" + global.CONFIG.App.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.LOG.Fatal("listen: %s\n", zap.Error(err))
		}
	}()
	// 优雅Shutdown（或重启）服务
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
