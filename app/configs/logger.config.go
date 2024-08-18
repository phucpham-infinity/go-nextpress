package configs

import (
	"github.com/phucpham-infinity/go-nextpress/app/global"
	"github.com/phucpham-infinity/go-nextpress/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.AppLogger(global.Config.Logger)
}
