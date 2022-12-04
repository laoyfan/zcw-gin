package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"time"
)

type Config struct {
	Disable      bool
	Driver       string
	Host         string
	Port         string
	Database     string
	Username     string
	Password     string
	Write        []Node
	Read         []Node
	Charset      string
	Collation    string
	MaxIdLeConns int
	MaxOpenConns int
	LogLevel     string
	Log          bool
}

type Node struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func Mysql(c Config) *gorm.DB {
	var (
		sources  []gorm.Dialector
		replicas []gorm.Dialector
	)

	main := mysql.New(mysql.Config{
		DSN:                       c.Username + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Database,
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	})

	if len(c.Write) > 0 {
		for _, node := range c.Write {
			sources = append(sources, Dialectic(node))
		}
	}

	if len(c.Read) > 0 {
		for _, node := range c.Read {
			replicas = append(replicas, Dialectic(node))
		}
	}

	db, err := gorm.Open(main, getConfig(c.LogLevel))
	if err != nil {
		fmt.Println("数据库连接异常:", err)
		return nil
	} else {
		err = db.Use(
			dbresolver.Register(
				dbresolver.Config{
					Sources:  sources,
					Replicas: replicas,
					Policy:   dbresolver.RandomPolicy{},
				}).
				SetConnMaxIdleTime(time.Hour).
				SetConnMaxLifetime(24 * time.Hour).
				SetMaxIdleConns(100).
				SetMaxOpenConns(200))
		if err != nil {
			fmt.Println("数据库配置异常:", err)
			return nil
		}
		return db
	}

}

func Dialectic(n Node) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN:                       n.Username + ":" + n.Password + "@tcp(" + n.Host + ":" + n.Port + ")/" + n.Database,
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	})
}

func getConfig(level string) *gorm.Config {
	c := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	l := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Microsecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	switch level {
	case "silent":
		c.Logger = l.LogMode(logger.Silent)
	case "error":
		c.Logger = l.LogMode(logger.Error)
	case "warn":
		c.Logger = l.LogMode(logger.Warn)
	case "info":
		c.Logger = l.LogMode(logger.Info)
	default:
		c.Logger = l.LogMode(logger.Info)
	}
	return c
}
