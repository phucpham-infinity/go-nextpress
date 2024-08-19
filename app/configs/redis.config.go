package configs

import (
	_context "context"
	"fmt"

	"github.com/phucpham-infinity/go-nextpress/app/context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = _context.Background()

func InitRedis() {
	logger := context.AppContext().GetLogger()
	config := context.AppContext().GetConfig()

	addr := "%s:%v"
	addr = fmt.Sprintf(addr, config.Redis.Host, config.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10,
	})

	if rdb != nil {
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			logger.Error("Error pinging redis", zap.Error(err))
			panic(err)
		}
	} else {
		logger.Error("Redis client is nil")
		panic("Redis client is nil")
	}
	context.AppContext().SetRedisClient(rdb)
	logger.Info("Connected to redis! " + addr)
}
