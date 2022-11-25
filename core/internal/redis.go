package internal

import (
	"zcw-gin/global"
	"zcw-gin/pkg/redis"
)

// 初始化redis

func Redis() {
	for _, info := range global.CONFIG.Redis {
		if info.Disable {
			continue
		}
		global.REDIS[info.Name] = redis.NewRedisClient(info)
		global.LOG.Info("redis:" + info.Name + "连接成功")
	}
}

// 关闭redis连接

func RedisClose() {
	if len(global.REDIS) > 0 {
		for name, client := range global.REDIS {
			if client != nil {
				_ = client.Close()
				global.LOG.Info("redis:" + name + "关闭连接")
			}
		}
	}
}
