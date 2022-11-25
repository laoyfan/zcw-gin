package internal

import (
	"zcw-gin/global"
	"zcw-gin/pkg/mysql"
)

// 初始化mysql

func Mysql() {
	for _, info := range global.CONFIG.Mysql {
		if info.Disable {
			continue
		}
		global.DB[info.Name] = mysql.NewMysql(info)
		global.LOG.Info("mysql:" + info.Name + "连接成功")
	}
}

// 关闭mysql连接

func MysqlClose() {
	if len(global.DB) > 0 {
		for name, db := range global.DB {
			if db != nil {
				s, _ := db.DB()
				_ = s.Close()
				global.LOG.Info("mysql:" + name + "关闭连接")
			}
		}
	}
}
