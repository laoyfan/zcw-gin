package internal

import (
	"fmt"
	"zcw-gin/pkg/database"

	"gorm.io/gorm"

	"zcw-gin/core"
)

func Database(name ...string) *gorm.DB {
	var (
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

		var (
			write []database.Node
			read  []database.Node
		)

		err := config.UnmarshalKey(dbname+".write", &write)
		if err != nil {
			write = []database.Node{}
		}
		err = config.UnmarshalKey(dbname+".read", &read)
		if err != nil {
			read = []database.Node{}
		}

		return database.Mysql(database.Config{
			Disable:      config.GetBool(dbname + ".disable"),
			Driver:       config.GetString(dbname + ".driver"),
			Host:         config.GetString(dbname + ".host"),
			Port:         config.GetString(dbname + ".port"),
			Database:     config.GetString(dbname + ".database"),
			Username:     config.GetString(dbname + ".username"),
			Password:     config.GetString(dbname + ".password"),
			Write:        write,
			Read:         read,
			Charset:      config.GetString(dbname + ".charset"),
			Collation:    config.GetString(dbname + ".collation"),
			MaxIdLeConns: config.GetInt(dbname + ".maxIdleConns"),
			MaxOpenConns: config.GetInt(dbname + ".maxOpenConns"),
			LogLevel:     config.GetString(dbname + ".logLevel"),
			Log:          config.GetBool(dbname + ".log"),
		})
	}).(*gorm.DB)
}
