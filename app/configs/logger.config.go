package configs

import (
	"github.com/phucpham-infinity/go-nextpress/app/context"
	"github.com/phucpham-infinity/go-nextpress/pkg/logger"
)

func InitLogger() {
	config := context.AppContext().GetConfig()
	context.AppContext().SetLogger(logger.AppLogger(config.Logger))
}
