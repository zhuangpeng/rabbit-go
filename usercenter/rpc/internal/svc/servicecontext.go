package svc

import (
	"backend/rabbit-go/usercenter/rpc/internal/config"
	"backend/rabbit-go/usercenter/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
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
