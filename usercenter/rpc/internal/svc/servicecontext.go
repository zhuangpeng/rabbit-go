package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/config"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/model"
)

type ServiceContext struct {
	Config config.Config
	RedisClient *redis.Redis
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(c.DatabaseConf.NewSqlConn(), c.Cache),
	}
}
