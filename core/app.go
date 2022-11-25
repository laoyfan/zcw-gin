package core

import (
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

func Config(name ...string) *internal.Config {
	return internal.NewConfig(name...)
}
