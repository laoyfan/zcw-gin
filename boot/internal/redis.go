package internal

import (
	"zcw-admin-server/global"
	"zcw-admin-server/pkg/redis"
)

// 初始化redis

func Redis() {
	for _, info := range global.CONFIG.Redis {
		if info.Disable {
			continue
		}
		global.REDIS[info.Name] = redis.NewRedisClient(info)
	}
}
