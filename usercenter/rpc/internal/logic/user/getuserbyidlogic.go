package userlogic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/zhuangpeng/rabbit-go/pkg/i18n"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/model"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *pb.UUIDReq) (UserInfoResp *pb.UserInfoResp, err error) {
	u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		logx.Errorw(err.Error(), logx.Field("query user when soft delete has err, user id is:", u.Id))
		return nil, statuserr.NewInvalidArgumentError(i18n.DatabaseError)
	}
	if err == model.ErrNotFound {
		logx.Errorw(err.Error(), logx.Field("soft delete user not found, user id is:", u.Id))
		return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
	}
	
	return &pb.UserInfoResp{}, nil
}