package global

import (
	"github.com/didip/tollbooth/limiter"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局变量

var (
	DB      = map[string]*gorm.DB{}
	REDIS   = map[string]*redis.Client{}
	LOG     *zap.Logger
	CONFIG  Config
	LIMITER *limiter.Limiter
	Trans   ut.Translator
	Golbal  = NewContainer()
)
