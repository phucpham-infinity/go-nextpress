package context

import (
	"github.com/phucpham-infinity/go-nextpress/app/common"
	"github.com/phucpham-infinity/go-nextpress/pkg/logger"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type IAppContext interface {
	GetMySqlConnection() *gorm.DB
	SetMySqlConnection(*gorm.DB) *appContext

	SetRedisClient(*redis.Client) *appContext
	GetRedisClient() *redis.Client

	SetLogger(*logger.LoggerZap) *appContext
	GetLogger() *logger.LoggerZap

	SetConfig(*common.Config) *appContext
	GetConfig() *common.Config
}

type appContext struct {
	DB     *gorm.DB
	RDB    *redis.Client
	Logger *logger.LoggerZap
	Config *common.Config
}

var appContextInstance *appContext

func AppContext() IAppContext {
	if appContextInstance != nil {
		return appContextInstance
	}
	appContextInstance = &appContext{
		DB:     nil,
		RDB:    nil,
		Logger: nil,
		Config: nil,
	}
	return appContextInstance
}

func (ctx *appContext) GetMySqlConnection() *gorm.DB {
	return appContextInstance.DB
}

func (ctx *appContext) SetMySqlConnection(DB *gorm.DB) *appContext {
	appContextInstance.DB = DB
	return appContextInstance
}

func (ctx *appContext) SetRedisClient(Client *redis.Client) *appContext {
	appContextInstance.RDB = Client
	return appContextInstance
}

func (ctx *appContext) GetRedisClient() *redis.Client {
	return appContextInstance.RDB
}

func (ctx *appContext) SetLogger(Logger *logger.LoggerZap) *appContext {
	appContextInstance.Logger = Logger
	return appContextInstance
}

func (ctx *appContext) GetLogger() *logger.LoggerZap {
	return appContextInstance.Logger
}

func (ctx *appContext) SetConfig(Config *common.Config) *appContext {
	appContextInstance.Config = Config
	return appContextInstance
}

func (ctx *appContext) GetConfig() *common.Config {
	return appContextInstance.Config
}
