package basic

import (
	"gorm.io/gorm"
	"zcw-admin-server/global"
)

type BaseModel struct{}

func (m *BaseModel) DB() *gorm.DB {
	return global.DB["basic"]
}
