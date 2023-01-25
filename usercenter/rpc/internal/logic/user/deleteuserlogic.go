package userlogic

import (
	"context"

	"backend/rabbit-go/usercenter/rpc/internal/svc"
	"backend/rabbit-go/usercenter/rpc/pb"

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
	// todo: add your logic here and delete this line

	return &pb.BaseResp{}, nil
}
