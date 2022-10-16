package boot

func init() {
	Viper()
	restart()
}

func restart() {
	Zap()
	initMysql()
	initRedis()
}
