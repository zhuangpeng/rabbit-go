// Code generated by goctl. DO NOT EDIT!
// Source: usercenter.proto

package server

import (
	"context"

	"backend/rabbit-go/usercenter/rpc/internal/logic/tenant"
	"backend/rabbit-go/usercenter/rpc/internal/svc"
	"backend/rabbit-go/usercenter/rpc/pb"
)

type TenantServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTenantServer
}

func NewTenantServer(svcCtx *svc.ServiceContext) *TenantServer {
	return &TenantServer{
		svcCtx: svcCtx,
	}
}

func (s *TenantServer) CreateOrUpdateTenant(ctx context.Context, in *pb.CreateOrUpdateTenantReq) (*pb.IDResp, error) {
	l := tenantlogic.NewCreateOrUpdateTenantLogic(ctx, s.svcCtx)
	return l.CreateOrUpdateTenant(in)
}

func (s *TenantServer) DeleteTenant(ctx context.Context, in *pb.UUIDReq) (*pb.BaseResp, error) {
	l := tenantlogic.NewDeleteTenantLogic(ctx, s.svcCtx)
	return l.DeleteTenant(in)
}

func (s *TenantServer) GetTenantTreeList(ctx context.Context, in *pb.TenantListReq) (*pb.TenantListResp, error) {
	l := tenantlogic.NewGetTenantTreeListLogic(ctx, s.svcCtx)
	return l.GetTenantTreeList(in)
}