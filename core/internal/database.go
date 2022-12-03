package internal

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"zcw-gin/core"
)

func Database(name ...string) *gorm.DB {
	var (
		ctx    = context.Background()
		prefix = "database"
		dbname = "default"
	)

	if len(name) > 0 && name[0] != "" {
		dbname = name[0]
	}

	dbKey := fmt.Sprintf("%s.%s", prefix, dbname)

	return core.Container.GetOrSetFunc(dbKey, func() interface{} {
		config := Config("config")
		disable := config.GetBool(dbname + ".disable")
		if disable {
			return nil
		}

		ctx.Err()
		return nil
		//return mysql.NewMysql()
	}).(*gorm.DB)
}
