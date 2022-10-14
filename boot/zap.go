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
	"zcw-admin-server/core"
	"zcw-admin-server/utils"
)

func Zap() {
	dir := ConfigMap["zap"]["director"]
	if ok, _ := utils.PathExists(dir); !ok {
		fmt.Println("创建日志文件夹", dir)
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Println("创建日志文件夹失败", err)
		}
	}

	core.Log = zap.New(zapcore.NewTee(getCores()...))
	showLine, _ := strconv.ParseBool(ConfigMap["zap"]["showLine"])
	if showLine {
		core.Log = core.Log.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(core.Log)
}

func getCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for zLevel := getLevel(); zLevel <= zapcore.FatalLevel; zLevel++ {
		cores = append(cores, getEncoderCore(zLevel, getLevelPriority(zLevel)))
	}
	return cores
}

// 获取配置对应level
func getLevel() zapcore.Level {
	switch ConfigMap["zap"]["level"] {
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

func getEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := getWriteSyncer(l.String())
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

func getWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	maxAge, _ := strconv.Atoi(ConfigMap["zap"]["maxAge"])
	fileWriter, err := rotatelogs.New(
		path.Join(ConfigMap["zap"]["director"], "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	logInConsole, _ := strconv.ParseBool(ConfigMap["zap"]["logInConsole"])
	if logInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

func getEncoder() zapcore.Encoder {
	format := ConfigMap["zap"]["format"]
	if format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  ConfigMap["zap"]["stackTraceKey"],
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    getEncodeLevel(),
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format(ConfigMap["zap"]["prefix"] + "2006/01/02 - 15:04:05.000"))
}

func getEncodeLevel() zapcore.LevelEncoder {
	encodeLevel := ConfigMap["zap"]["encodeLevel"]
	switch {
	case encodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case encodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case encodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case encodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func getLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}
