package internal

import (
	"fmt"
	"github.com/spf13/viper"

	"zcw-gin/core"
)

// Config 获取config目录下不同文件配置, name为文件名,默认为app
func Config(name ...string) *viper.Viper {

	var (
		prefix = "config"
		key    = "app"
	)

	if len(name) > 0 && name[0] != "" {
		key = name[0]
	}

	configKey := fmt.Sprintf("%s.%s", prefix, key)
	return core.Container.GetOrSetFunc(configKey, func() interface{} {
		v := viper.New()
		v.AddConfigPath("config")
		v.SetConfigName(key)
		err := v.ReadInConfig()
		if err != nil {
			return nil
		}
		v.WatchConfig()
		return v
	}).(*viper.Viper)
}
