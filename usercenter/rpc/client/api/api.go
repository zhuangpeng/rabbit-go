// Code generated by goctl. DO NOT EDIT!
// Source: usercenter.proto

package client

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiInfo               = pb.ApiInfo
	BaseResp              = pb.BaseResp
	BatchCreateUserReq    = pb.BatchCreateUserReq
	ChangeDept            = pb.ChangeDept
	ChangePasswdReq       = pb.ChangePasswdReq
	ChangePositionReq     = pb.ChangePositionReq
	CreateApiReq          = pb.CreateApiReq
	CreateMenuParamReq    = pb.CreateMenuParamReq
	CreateMenuReq         = pb.CreateMenuReq
	CreateRoleReq         = pb.CreateRoleReq
	CreateStationReq      = pb.CreateStationReq
	CreateTenantReq       = pb.CreateTenantReq
	CreateTokenReq        = pb.CreateTokenReq
	CreateUserReq         = pb.CreateUserReq
	Empty                 = pb.Empty
	GetApiListResp        = pb.GetApiListResp
	GetApiResp            = pb.GetApiResp
	GetMenuListReq        = pb.GetMenuListReq
	GetMenuListResp       = pb.GetMenuListResp
	GetParamsByMenuResp   = pb.GetParamsByMenuResp
	GetRoleResp           = pb.GetRoleResp
	GetStationByUserResp  = pb.GetStationByUserResp
	GetStationListResp    = pb.GetStationListResp
	GetStationResp        = pb.GetStationResp
	GetTeantListResp      = pb.GetTeantListResp
	GetTenantResp         = pb.GetTenantResp
	GetTokenListReq       = pb.GetTokenListReq
	GetTokenListResp      = pb.GetTokenListResp
	GetTokenResp          = pb.GetTokenResp
	GrantApiToRoleReq     = pb.GrantApiToRoleReq
	GrantMenuToRoleReq    = pb.GrantMenuToRoleReq
	GrantMenusToRoleReq   = pb.GrantMenusToRoleReq
	GrantRoleReq          = pb.GrantRoleReq
	GrantRoleToStationReq = pb.GrantRoleToStationReq
	IDReq                 = pb.IDReq
	IDResp                = pb.IDResp
	IDsReq                = pb.IDsReq
	IDsResp               = pb.IDsResp
	LoginReq              = pb.LoginReq
	LoginResp             = pb.LoginResp
	MenuInfo              = pb.MenuInfo
	MenuParamInfo         = pb.MenuParamInfo
	Meta                  = pb.Meta
	PageInfoReq           = pb.PageInfoReq
	ReplaceUserReq        = pb.ReplaceUserReq
	RoleInfo              = pb.RoleInfo
	RoleListResp          = pb.RoleListResp
	StationInfo           = pb.StationInfo
	StationUserReq        = pb.StationUserReq
	StatusCodeReq         = pb.StatusCodeReq
	TSRU                  = pb.TSRU
	TenantInfo            = pb.TenantInfo
	TokenInfo             = pb.TokenInfo
	UpdateApiReq          = pb.UpdateApiReq
	UpdateMenuParamsReq   = pb.UpdateMenuParamsReq
	UpdateMenuReq         = pb.UpdateMenuReq
	UpdatePeriodReq       = pb.UpdatePeriodReq
	UpdateProfileReq      = pb.UpdateProfileReq
	UpdateRoleReq         = pb.UpdateRoleReq
	UpdateStationReq      = pb.UpdateStationReq
	UpdateStatusReq       = pb.UpdateStatusReq
	UpdateTenantReq       = pb.UpdateTenantReq
	UserInfoResp          = pb.UserInfoResp
	UserListResp          = pb.UserListResp

	Api interface {
		CreateApi(ctx context.Context, in *CreateApiReq, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteApi(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		UpdateApi(ctx context.Context, in *UpdateApiReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetApi(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*GetApiResp, error)
		GetApiListByMenu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*GetApiListResp, error)
	}

	defaultApi struct {
		cli zrpc.Client
	}
)

func NewApi(cli zrpc.Client) Api {
	return &defaultApi{
		cli: cli,
	}
}

func (m *defaultApi) CreateApi(ctx context.Context, in *CreateApiReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewApiClient(m.cli.Conn())
	return client.CreateApi(ctx, in, opts...)
}

func (m *defaultApi) DeleteApi(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewApiClient(m.cli.Conn())
	return client.DeleteApi(ctx, in, opts...)
}

func (m *defaultApi) UpdateApi(ctx context.Context, in *UpdateApiReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewApiClient(m.cli.Conn())
	return client.UpdateApi(ctx, in, opts...)
}

func (m *defaultApi) GetApi(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*GetApiResp, error) {
	client := pb.NewApiClient(m.cli.Conn())
	return client.GetApi(ctx, in, opts...)
}

func (m *defaultApi) GetApiListByMenu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*GetApiListResp, error) {
	client := pb.NewApiClient(m.cli.Conn())
	return client.GetApiListByMenu(ctx, in, opts...)
}
