package main

import (
	"flag"
	"fmt"

	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/config"
	apiServer "github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/server/api"
	menuServer "github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/server/menu"
	roleServer "github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/server/role"
	stationServer "github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/server/station"
	tenantServer "github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/server/tenant"
	tokenServer "github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/server/token"
	userServer "github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/server/user"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		pb.RegisterStationServer(grpcServer, stationServer.NewStationServer(ctx))
		pb.RegisterRoleServer(grpcServer, roleServer.NewRoleServer(ctx))
		pb.RegisterMenuServer(grpcServer, menuServer.NewMenuServer(ctx))
		pb.RegisterTenantServer(grpcServer, tenantServer.NewTenantServer(ctx))
		pb.RegisterApiServer(grpcServer, apiServer.NewApiServer(ctx))
		pb.RegisterTokenServer(grpcServer, tokenServer.NewTokenServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
