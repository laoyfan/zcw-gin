package global

import (
	"github.com/didip/tollbooth/limiter"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 当前文件挂载全局

var (
	DB      = map[string]*gorm.DB{}
	REDIS   = map[string]*redis.Client{}
	LOG     *zap.Logger
	CONFIG  Config
	LIMITER *limiter.Limiter
)
