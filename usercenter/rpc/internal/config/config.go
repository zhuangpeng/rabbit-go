package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhuangpeng/rabbit-go/pkg/config"
)

type Config struct {
	zrpc.RpcServerConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	DatabaseConf config.DatabaseConf
	RedisConf    redis.RedisConf
	Cache 		 cache.CacheConf
	CasbinConf   config.CasbinConf
}
