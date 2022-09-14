package boot

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"zcw-admin-server/config/entity"
	"zcw-admin-server/utils"
)

func Viper() {
	// 读取env
	env := viper.New()                     // 创建env实例
	env.SetConfigFile(".env")              // 读取指定文件
	envMap := make(map[string]interface{}) // 接受env数据map
	if err := env.ReadInConfig(); err == nil {
		env.WatchConfig() // 监听env
		env.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file change:", e.Name) // @todo::后续日志模块写入
			if err = env.Unmarshal(&envMap); err != nil {
				fmt.Println(err)
			}
		})
		if err = env.Unmarshal(&envMap); err != nil {
			fmt.Println(err)
		}

	}

	// 读取config
	fileNames := utils.GetPathFileNames("config")
	if len(fileNames) > 0 {
		for _, fileName := range fileNames {
			v := viper.New()
			v.SetConfigFile("config/" + fileName)
			if err := v.ReadInConfig(); err != nil {
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}
			v.WatchConfig()
			v.OnConfigChange(func(e fsnotify.Event) {
				fmt.Println("config file change:", e.Name)
				if err := v.Unmarshal(&entity.Config{}); err != nil {
					fmt.Println(err)
				}
			})
			if err := v.Unmarshal(&entity.CONFIG); err != nil {
				fmt.Println(err)
			}
		}

	}
}
