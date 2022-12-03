package core

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"zcw-gin/core/internal"
)

var (
	// Container 创建全局容器
	Container = internal.NewContainer()
)

func DB(name ...string) *gorm.DB {
	return internal.Database(name...)
}

func Config(name ...string) *viper.Viper {
	return internal.Config(name...)
}
