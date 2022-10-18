package internal

import (
	"zcw-admin-server/global"
	"zcw-admin-server/pkg/mysql"
)

// 初始化mysql

func Mysql() {
	for _, info := range global.CONFIG.Mysql {
		if info.Disable {
			continue
		}
		global.DB[info.Name] = mysql.NewMysql(info)
	}
}
