package userlogic

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchCreateUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchCreateUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchCreateUsersLogic {
	return &BatchCreateUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchCreateUsersLogic) BatchCreateUsers(in *pb.BatchCreateUserReq) (*pb.IDsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.IDsResp{}, nil
}
