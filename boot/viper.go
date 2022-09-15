package boot

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"zcw-admin-server/utils"
)

const (
	ConfigDir = "config" //配置文件夹
)

var (
	EnvMap    = make(map[string]string)            // 接受env数据map
	ConfigMap = make(map[string]map[string]string) // 接受config数据
)

// Viper 自动读取配置
// 支持根目录.env文件热更新
// 多配置读取
func Viper() {
	// 读取ConfigDir目录下所有配置
	fileNames := utils.GetPathFileNames(ConfigDir) // 获取config文件夹下配置文件名称
	if len(fileNames) > 0 {
		for _, fileName := range fileNames {
			v := viper.New()                            // 每个文件定义一个实例
			v.SetConfigFile(ConfigDir + "/" + fileName) // 配置文件
			if err := v.ReadInConfig(); err != nil {    // 读取文件 此时异常需要panic
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}
			if err := v.Unmarshal(&ConfigMap); err != nil { // 配置写入
				fmt.Println(err) //@todo::后续日志模块写入
			}
		}
	}
	// 读取env
	env := viper.New()                         // 创建env实例
	env.SetConfigFile(".env")                  // 读取指定文件
	if err := env.ReadInConfig(); err == nil { // 存在.env文件则处理
		env.WatchConfig() // 监听env
		env.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("env file change:", e.Name)       // @todo::后续日志模块写入
			if err = env.Unmarshal(&EnvMap); err != nil { // 无异常处理
				fmt.Println(err) //@todo::后续日志模块写入
			}
			getEnv()
		})
		if err = env.Unmarshal(&EnvMap); err != nil { // 无异常处理
			fmt.Println(err) //@todo::后续日志模块写入
		}
	} else {
		fmt.Println(err) //@todo::后续日志模块写入
	}
	getEnv()
}

// 写入env文件数据
func getEnv() {
	for cKey, singleMap := range ConfigMap { // 循环总配置
		for sKey, _ := range singleMap { // 循环子配置
			eKey := cKey + "_" + sKey           // 组装env的key
			if eValue, ok := EnvMap[eKey]; ok { // 判断env是否有相关数据
				ConfigMap[cKey][sKey] = eValue
			}
		}
	}
}
