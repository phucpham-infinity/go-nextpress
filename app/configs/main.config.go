package configs

import "github.com/phucpham-infinity/go-nextpress/app/context"

func Run() {
	InitEnv()
	InitLogger()
	InitMysql()
	InitRedis()

	logger := context.AppContext().GetLogger()
	logger.Info("Configs loaded!")
}
