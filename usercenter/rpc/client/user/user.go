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

	User interface {
		CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*BaseResp, error)
		BatchCreateUsers(ctx context.Context, in *BatchCreateUserReq, opts ...grpc.CallOption) (*IDsResp, error)
		DeleteUser(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		BatchDeleteUser(ctx context.Context, in *IDsResp, opts ...grpc.CallOption) (*BaseResp, error)
		UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetUser(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		GetUserList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*UserListResp, error)
		UpdateUserStatus(ctx context.Context, in *StatusCodeReq, opts ...grpc.CallOption) (*BaseResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		ChangePassword(ctx context.Context, in *ChangePasswdReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultUser) BatchCreateUsers(ctx context.Context, in *BatchCreateUserReq, opts ...grpc.CallOption) (*IDsResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.BatchCreateUsers(ctx, in, opts...)
}

func (m *defaultUser) DeleteUser(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

func (m *defaultUser) BatchDeleteUser(ctx context.Context, in *IDsResp, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.BatchDeleteUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UpdateProfile(ctx, in, opts...)
}

func (m *defaultUser) GetUser(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) GetUserList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUserList(ctx, in, opts...)
}

func (m *defaultUser) UpdateUserStatus(ctx context.Context, in *StatusCodeReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UpdateUserStatus(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) ChangePassword(ctx context.Context, in *ChangePasswdReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.ChangePassword(ctx, in, opts...)
}
