package global

import (
	"github.com/phucpham-infinity/go-nextpress/app/structs"
	"github.com/phucpham-infinity/go-nextpress/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config structs.Config
	Logger *logger.LoggerZap
	DB     *gorm.DB
	RDB    *redis.Client
)
