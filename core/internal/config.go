package internal

import (
	"fmt"
	"io/ioutil"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"zcw-gin/core"
)

const ConfigDir = "config" //配置文件夹

func NewConfig(name ...string) *viper.Viper {

	if len(name) > 0 && name[0] != "" {
		configKey := fmt.Sprintf("%s.%s", ConfigDir, name[0])
		return core.Container.GetOrSetFunc(configKey, func() interface{} {
			v := viper.New()
			err := v.ReadInConfig()
			if err != nil {
				return nil
			}
			v.WatchConfig()
			return v
		}).(*viper.Viper)
	}

	files, _ := ioutil.ReadDir(ConfigDir)
	for _, file := range files {
		if !file.IsDir() {
			names = append(names, file.Name())
		}
	}

	return core.Container.GetOrSetFunc(configKey, func() interface{} {
		v := viper.New()                         // 每个文件定义一个实例
		v.SetConfigFile(ConfigDir + "/" + path)  // 配置文件
		if err := v.ReadInConfig(); err != nil { // 读取文件 此时异常需要panic
			panic(fmt.Errorf("读取配置文件异常: %s \n", err))
		}
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("配置文件变更:", e.Name)
			if err := v.Unmarshal(&Config{}); err != nil { // 配置写入
				panic(fmt.Errorf("Config转换异常: %s \n", err))
			}
		})
		if err := v.Unmarshal(&Config{}); err != nil { // 配置写入
			panic(fmt.Errorf("Config转换异常: %s \n", err))
		}
		return Config{}
	}).(*Config)
}
