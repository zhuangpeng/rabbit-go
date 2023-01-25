// Code generated by goctl. DO NOT EDIT!
// Source: usercenter.proto

package client

import (
	"context"

	"backend/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiInfo                    = pb.ApiInfo
	ApiListResp                = pb.ApiListResp
	ApiPageReq                 = pb.ApiPageReq
	BaseResp                   = pb.BaseResp
	CallbackReq                = pb.CallbackReq
	ChangePasswordReq          = pb.ChangePasswordReq
	CreateOrUpdateMenuParamReq = pb.CreateOrUpdateMenuParamReq
	CreateOrUpdateMenuReq      = pb.CreateOrUpdateMenuReq
	CreateOrUpdateTenantReq    = pb.CreateOrUpdateTenantReq
	CreateOrUpdateUserReq      = pb.CreateOrUpdateUserReq
	CreatePolicyReq            = pb.CreatePolicyReq
	DictionaryDetail           = pb.DictionaryDetail
	DictionaryDetailList       = pb.DictionaryDetailList
	DictionaryDetailReq        = pb.DictionaryDetailReq
	DictionaryInfo             = pb.DictionaryInfo
	DictionaryList             = pb.DictionaryList
	DictionaryPageReq          = pb.DictionaryPageReq
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
	OauthLoginReq              = pb.OauthLoginReq
	OauthRedirectResp          = pb.OauthRedirectResp
	PageInfoReq                = pb.PageInfoReq
	PolicyPartInfo             = pb.PolicyPartInfo
	ProviderInfo               = pb.ProviderInfo
	ProviderListResp           = pb.ProviderListResp
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
	TokenInfo                  = pb.TokenInfo
	TokenListReq               = pb.TokenListReq
	TokenListResp              = pb.TokenListResp
	UUIDReq                    = pb.UUIDReq
	UUIDsReq                   = pb.UUIDsReq
	UpdatePolicyReq            = pb.UpdatePolicyReq
	UpdateProfileReq           = pb.UpdateProfileReq
	UserInfoResp               = pb.UserInfoResp
	UserListResp               = pb.UserListResp

	Api interface {
		CreateOrUpdateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteApi(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetApiList(ctx context.Context, in *ApiPageReq, opts ...grpc.CallOption) (*ApiListResp, error)
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

func (m *defaultApi) CreateOrUpdateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewApiClient(m.cli.Conn())
	return client.CreateOrUpdateApi(ctx, in, opts...)
}

func (m *defaultApi) DeleteApi(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pb.NewApiClient(m.cli.Conn())
	return client.DeleteApi(ctx, in, opts...)
}

func (m *defaultApi) GetApiList(ctx context.Context, in *ApiPageReq, opts ...grpc.CallOption) (*ApiListResp, error) {
	client := pb.NewApiClient(m.cli.Conn())
	return client.GetApiList(ctx, in, opts...)
}