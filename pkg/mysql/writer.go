package mysql

import (
	"fmt"
	"gorm.io/gorm/logger"
	"zcw-admin-server/global"
)

type Writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *Writer {
	return &Writer{w}
}

func (w *Writer) Printf(message string, data ...interface{}) {
	w.Writer.Printf(message, data...)
	if global.CONFIG.App.Mode == "debug" {
		w.Writer.Printf(message, data...)
	} else {
		global.LOG.Info(fmt.Sprintf(message+"\n", data))
	}
}
