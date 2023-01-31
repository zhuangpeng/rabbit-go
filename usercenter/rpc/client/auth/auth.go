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
	ApiInfo                    = pb.ApiInfo
	ApiListResp                = pb.ApiListResp
	ApiPageReq                 = pb.ApiPageReq
	BaseResp                   = pb.BaseResp
	ChangePasswordReq          = pb.ChangePasswordReq
	CreateOrUpdateMenuParamReq = pb.CreateOrUpdateMenuParamReq
	CreateOrUpdateMenuReq      = pb.CreateOrUpdateMenuReq
	CreateOrUpdateTenantReq    = pb.CreateOrUpdateTenantReq
	CreateOrUpdateUserReq      = pb.CreateOrUpdateUserReq
	CreatePolicyReq            = pb.CreatePolicyReq
	Empty                      = pb.Empty
	GetUserListReq             = pb.GetUserListReq
	IDReq                      = pb.IDReq
	IDResp                     = pb.IDResp
	IDsReq                     = pb.IDsReq
	LoginReq                   = pb.LoginReq
	LoginResp                  = pb.LoginResp
	MenuInfo                   = pb.MenuInfo
	MenuInfoList               = pb.MenuInfoList
	MenuParamListResp          = pb.MenuParamListResp
	MenuParamResp              = pb.MenuParamResp
	MenuRoleInfo               = pb.MenuRoleInfo
	MenuRoleListResp           = pb.MenuRoleListResp
	Meta                       = pb.Meta
	PageInfoReq                = pb.PageInfoReq
	PolicyPartInfo             = pb.PolicyPartInfo
	RoleInfo                   = pb.RoleInfo
	RoleListResp               = pb.RoleListResp
	RoleMenuAuthorityReq       = pb.RoleMenuAuthorityReq
	RoleMenuAuthorityResp      = pb.RoleMenuAuthorityResp
	StatusCodeReq              = pb.StatusCodeReq
	StatusCodeUUIDReq          = pb.StatusCodeUUIDReq
	TenantInfo                 = pb.TenantInfo
	TenantListReq              = pb.TenantListReq
	TenantListResp             = pb.TenantListResp
	TenantTreeInfo             = pb.TenantTreeInfo
	TenantTreeListResp         = pb.TenantTreeListResp
	UUIDReq                    = pb.UUIDReq
	UUIDsReq                   = pb.UUIDsReq
	UpdatePolicyReq            = pb.UpdatePolicyReq
	UpdateProfileReq           = pb.UpdateProfileReq
	UserInfoResp               = pb.UserInfoResp
	UserListResp               = pb.UserListResp

	Auth interface {
		GetMenuAuthority(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleMenuAuthorityResp, error)
		CreateOrUpdateMenuAuthority(ctx context.Context, in *RoleMenuAuthorityReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultAuth struct {
		cli zrpc.Client
	}
)

func NewAuth(cli zrpc.Client) Auth {
	return &defaultAuth{
		cli: cli,
	}
}

func (m *defaultAuth) GetMenuAuthority(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleMenuAuthorityResp, error) {
	client := pb.NewAuthClient(m.cli.Conn())
	return client.GetMenuAuthority(ctx, in, opts...)
}

func (m *defaultAuth) CreateOrUpdateMenuAuthority(ctx context.Context, in *RoleMenuAuthorityReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewAuthClient(m.cli.Conn())
	return client.CreateOrUpdateMenuAuthority(ctx, in, opts...)
}
