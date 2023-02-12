package svc

import (
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/config"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/model"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(c.DatabaseConf.NewSqlConn(), c.Cache),
	}
}
