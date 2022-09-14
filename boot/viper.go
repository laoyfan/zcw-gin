package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"zcw-admin-server/utils"
)

func Viper() {
	// 读取env
	env := viper.New()
	env.SetConfigFile(".env")
	if err := env.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(env.GetString("APP_NAME"), 111111111)
	// 读取config
	configs := utils.GetPathFileNames("config")
	if len(configs) > 0 {

	}
	fmt.Println(configs)

	v := viper.New()

	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	fmt.Println(gin.Mode())
	fmt.Println(v.GetString("name"))
}
