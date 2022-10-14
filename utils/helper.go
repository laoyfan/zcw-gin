package utils

import (
	"go.uber.org/zap"
	"zcw-admin-server/core"
)

func WriteLog(msg string, err error) {
	core.Log.Error(msg, zap.Error(err))
}
