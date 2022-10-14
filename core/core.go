package core

import (
	"go.uber.org/zap"
)

type config struct {
	Zap Zap `yaml:"zap"`
}

var (
	Config config
	Log    *zap.Logger
)

func init() {
	Viper()
	initZap()
}
