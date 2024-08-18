package configs

import "github.com/phucpham-infinity/go-nextpress/app/global"

func Run() {
	InitEnv()
	InitLogger()
	InitMysql()
	InitRedis()

	global.Logger.Info("Configs loaded!")
}
