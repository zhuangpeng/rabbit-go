package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhuangpeng/rabbit-go/pkg/i18n"
	"github.com/zhuangpeng/rabbit-go/usercenter/api/internal/config"
	client "github.com/zhuangpeng/rabbit-go/usercenter/rpc/client/user"
)

type ServiceContext struct {
	Config   config.Config
	Trans    *i18n.Translator
	UserRpc  client.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	trans := &i18n.Translator{}
	trans.NewBundle(i18n.LocaleFS)
	trans.NewTranslator()

	return &ServiceContext{
		Config: c,
		UserRpc: client.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		Trans: trans,
	}
}
