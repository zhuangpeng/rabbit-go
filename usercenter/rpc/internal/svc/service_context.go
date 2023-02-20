package svc

import (
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/config"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/model"
)

type ServiceContext struct {
	Config       config.Config
	UserModel    model.UserModel
	StationModel model.StationModel
	TenantModel  model.TenantModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config      : c,
		UserModel   : model.NewUserModel(c.DatabaseConf.NewSqlConn(), c.Cache),
		StationModel: model.NewStationModel(c.DatabaseConf.NewSqlConn(), c.Cache),
		TenantModel:  model.NewTenantModel(c.DatabaseConf.NewSqlConn(), c.Cache),
	}
}
