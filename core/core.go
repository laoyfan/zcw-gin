package core

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config config
	Log    *zap.Logger
	DB     map[string]*gorm.DB
)

type config struct {
	App   App     `yaml:"app"`
	Zap   Zap     `yaml:"zap"`
	Mysql []Mysql `yaml:"mysql"`
	Redis Redis   `yaml:"rides"`
}

type App struct {
	Env   string
	Debug bool
}

func init() {
	initViper()
	restart()
}

func restart() {
	initZap()
	initMysql()
}
