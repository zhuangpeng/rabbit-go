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

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *pb.UUIDReq) (*pb.BaseResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		logx.Errorw(err.Error(), logx.Field("query user when soft delete has err, user id is:", user.Id))
		return nil, statuserr.NewInvalidArgumentError(i18n.DatabaseError)
	}
	if err == model.ErrNotFound {
		logx.Errorw(err.Error(), logx.Field("soft delete user not found, user id is:", user.Id))
		return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
	}

	err = l.svcCtx.UserModel.DeleteSoft(l.ctx, nil, user)
	if err != nil {
		logx.Errorw(err.Error(), logx.Field("deleteSoft user has error: %+v", user))
		return nil, statuserr.NewInvalidArgumentError(i18n.DatabaseError)
	}	

	return &pb.BaseResp{
		Msg: i18n.Success,
	}, nil
}
