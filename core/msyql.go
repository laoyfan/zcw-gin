package core

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

type Mysql struct {
	Name         string
	Disable      bool
	Type         string
	Node         []Node
	MaxIdleConns int
	MaxOpenConns int
	LogLevel     string
	Log          bool
}

type Node struct {
	Path     string
	Port     string
	Database string
	Username string
	Password string
	Config   string
	Role     string
}

type Writer struct {
	logger.Writer
}

func initMysql() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range Config.Mysql {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.Name] = getMysqlDB(info)
		default:
			continue
		}

	}
	DB = dbMap
}

func getMysqlDB(m Mysql) *gorm.DB {
	var (
		sources  []gorm.Dialector
		replicas []gorm.Dialector
	)

	for _, node := range m.Node {
		dsn := node.Username + ":" + node.Password + "@tcp(" + node.Path + ":" + node.Port + ")/" + node.Database + "?" + node.Config

		mysqlConfig := mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         191,
			SkipInitializeWithVersion: false,
		}

		if node.Role == "slave" {
			replicas = append(replicas, mysql.New(mysqlConfig))
		} else {
			sources = append(sources, mysql.New(mysqlConfig))
		}
	}
	if db, err := gorm.Open(sources[0], getConfig(m)); err != nil {
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
			return nil
		}
		return db
	}
}

func getConfig(m Mysql) *gorm.Config {
	c := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	l := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Microsecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	switch m.LogLevel {
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

func NewWriter(w logger.Writer) *Writer {
	return &Writer{w}
}

func (w *Writer) Printf(message string, data ...interface{}) {
	w.Writer.Printf(message, data...)
	if Config.App.Debug {
		w.Writer.Printf(message, data...)
	} else {
		Log.Info(fmt.Sprintf(message+"\n", data))
	}
}
