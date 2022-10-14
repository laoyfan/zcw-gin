package boot

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
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
				panic(fmt.Errorf("读取配置文件异常: %s \n", err))
			}
			if err := v.Unmarshal(&ConfigMap); err != nil { // 配置写入
				fmt.Println(err)
			}
		}
	}
	// 读取env
	env := viper.New()                         // 创建env实例
	env.SetConfigFile(".env")                  // 读取指定文件
	if err := env.ReadInConfig(); err == nil { // 存在.env文件则处理
		env.WatchConfig() // 监听env
		env.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("env文件已变更:", e.Name)
			if err = env.Unmarshal(&EnvMap); err != nil { // 无异常处理
				fmt.Println(err)
			}
			getEnv()
		})
		if err = env.Unmarshal(&EnvMap); err != nil { // 无异常处理
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
	getEnv()
}

// 写入env文件数据
func getEnv() {
	for cKey, singleMap := range ConfigMap { // 循环总配置
		for sKey, sValue := range singleMap { // 循环子配置
			eKey := cKey + "_" + sKey           // 组装env的key
			if eValue, ok := EnvMap[eKey]; ok { // 判断env是否有相关数据
				ConfigMap[cKey][sKey] = eValue
			} else {
				sArr := strings.Split(sValue, "|")
				if len(sArr) > 1 {
					ConfigMap[cKey][sKey] = sArr[1]
				}
			}
		}
	}
}
