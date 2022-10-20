package internal

import (
	"github.com/didip/tollbooth"
	"zcw-admin-server/global"
)

// 创建限流器实例

func Limiter() {
	global.LIMITER = tollbooth.NewLimiter(global.CONFIG.App.Limit, nil)
}
