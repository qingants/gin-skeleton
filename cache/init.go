package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/qingants/gin-skeleton/pkg/xredis"
	"github.com/qingants/gin-skeleton/setting"
)

var (
	rdb *redis.Client
)

func Open() {
	rdb = xredis.Open(setting.RdbOpts)
}

func Close() {
	xredis.Close(rdb)
}
