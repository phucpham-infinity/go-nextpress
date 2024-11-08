package context

import (
	"database/sql"

	common_struct "github.com/phucpham-infinity/go-nextpress/app/common/struct"
	"github.com/phucpham-infinity/go-nextpress/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type IAppContext interface {
	GetMySqlConnection() *sql.DB
	SetMySqlConnection(*sql.DB) *appContext

	SetRedisClient(*redis.Client) *appContext
	GetRedisClient() *redis.Client

	SetLogger(*logger.LoggerZap) *appContext
	GetLogger() *logger.LoggerZap

	SetConfig(*common_struct.Config) *appContext
	GetConfig() *common_struct.Config
}

type appContext struct {
	DB     *sql.DB
	RDB    *redis.Client
	Logger *logger.LoggerZap
	Config *common_struct.Config
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

func (ctx *appContext) GetMySqlConnection() *sql.DB {
	return appContextInstance.DB
}

func (ctx *appContext) SetMySqlConnection(DB *sql.DB) *appContext {
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

func (ctx *appContext) SetConfig(Config *common_struct.Config) *appContext {
	appContextInstance.Config = Config
	return appContextInstance
}

func (ctx *appContext) GetConfig() *common_struct.Config {
	return appContextInstance.Config
}
