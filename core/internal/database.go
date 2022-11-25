package internal

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"zcw-gin/core"
	"zcw-gin/pkg/mysql"
)

func Database(name ...string) *gorm.DB {
	var (
		ctx    = context.Background()
		dbname = "default"
		prefix = "database"
	)

	if len(name) > 0 && name[0] != "" {
		dbname = name[0]
	}

	dbKey := fmt.Sprintf("%s.%s", prefix, dbname)

	db := core.Container.GetOrSetFunc(dbKey, func() interface{} {
		return mysql.NewMysql()
	})

	if db != nil {
		return db.(*gorm.DB)
	}
	return nil
}
