package boot

import (
	"github.com/go-redis/redis/v8"
	"zcw-admin-server/global"
	"zcw-admin-server/initialize"
)

func initRedis() {
	redisMap := make(map[string]*redis.Client)
	for _, info := range global.CONFIG.Redis {
		if info.Disable {
			continue
		}
		redisMap[info.Name] = initialize.NewRedisClient(info)
	}
	global.REDIS = redisMap
}
