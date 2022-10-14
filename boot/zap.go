package boot

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"strconv"
	"time"
	"zcw-admin-server/utils"
)

// 获取配置中参数
var (
	// 输出文件夹
	dir = "log"
	// 日志等级
	level  = ConfigMap["zap"]["level"]
	maxAge = 0
	format = "json"
)

type Zap struct {
}

// 初始化zap
func init() {
	if _, ok := ConfigMap["zap"]["director"]; ok {
		dir = ConfigMap["zap"]["director"]
	}
	if _, ok := ConfigMap["zap"]["level"]; ok {
		level = ConfigMap["zap"]["level"]
	}
	if _, ok := ConfigMap["zap"]["maxAge"]; ok {
		maxAge, _ = strconv.Atoi(ConfigMap["zap"]["maxAge"])
	}
}

func initZap() {

	if ok, _ := utils.PathExists(dir); !ok {
		fmt.Println("创建日志文件夹", dir)
		_ = os.Mkdir(dir, os.ModePerm)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

}

func cores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := getLevel(); level <= zapcore.FatalLevel; level++ {

		cores = append(cores)
	}
	return
}

func getLevel() zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func getWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(dir, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.GVA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

func getEncoder() zapcore.Encoder {
	if global.GVA_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}
