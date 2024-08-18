package configs

import (
	"context"
	"fmt"

	"github.com/phucpham-infinity/go-nextpress/app/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	addr := "%s:%v"
	addr = fmt.Sprintf(addr, global.Config.Redis.Host, global.Config.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10,
	})

	if rdb != nil {
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			global.Logger.Error("Error pinging redis", zap.Error(err))
			panic(err)
		}
	} else {
		global.Logger.Error("Redis client is nil")
		panic("Redis client is nil")
	}
	global.RDB = rdb
	global.Logger.Info("Connected to redis! " + addr)
}
