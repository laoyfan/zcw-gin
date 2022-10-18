package boot

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"zcw-admin-server/boot/internal"
	"zcw-admin-server/global"
	"zcw-admin-server/utils"
)

const ConfigDir = "config" //配置文件夹

func init() {
	// 读取ConfigDir目录下所有配置
	fileNames := utils.GetPathFileNames(ConfigDir) // 获取config文件夹下配置文件名称
	if len(fileNames) > 0 {
		for _, fileName := range fileNames {
			v := viper.New()                            // 每个文件定义一个实例
			v.SetConfigFile(ConfigDir + "/" + fileName) // 配置文件
			if err := v.ReadInConfig(); err != nil {    // 读取文件 此时异常需要panic
				panic(fmt.Errorf("读取配置文件异常: %s \n", err))
			}
			v.WatchConfig()
			v.OnConfigChange(func(e fsnotify.Event) {
				fmt.Println("配置文件变更:", e.Name)
				if err := v.Unmarshal(&global.CONFIG); err != nil { // 配置写入
					panic(fmt.Errorf("Config转换异常: %s \n", err))
				}
				created()
			})
			if err := v.Unmarshal(&global.CONFIG); err != nil { // 配置写入
				panic(fmt.Errorf("Config转换异常: %s \n", err))
			}
		}
	}
	created()
}

func created() {
	internal.Zap()
	internal.Mysql()
	internal.Redis()
}
