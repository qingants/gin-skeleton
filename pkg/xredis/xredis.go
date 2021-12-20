package xredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Open(opts *redis.Options) *redis.Client {
	client := redis.NewClient(opts)
	zap.S().Infof("connect to redis, addr: %s, pwd: %s", opts.Addr, opts.Password)
	if err := client.Ping(context.Background()).Err(); err != nil {
		zap.L().Fatal("Connect Redis: %s Error", zap.Error(err))
	}
	return client
}

func Close(client *redis.Client) {
	//client.FlushAll(context.Background())

	if err := client.Close(); err != nil {
		zap.L().Error("Close Redis Error", zap.String("addr", client.Options().Addr))
	}
}

