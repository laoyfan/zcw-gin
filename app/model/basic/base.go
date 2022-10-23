package basic

import (
	"gorm.io/gorm"
	"zcw-gin/global"
)

type BaseModel struct{}

func (m *BaseModel) DB() *gorm.DB {
	return global.DB["basic"]
}
