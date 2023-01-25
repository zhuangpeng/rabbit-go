package userlogic

import (
	"context"

	"backend/rabbit-go/usercenter/rpc/internal/svc"
	"backend/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateUserLogic {
	return &CreateOrUpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateUserLogic) CreateOrUpdateUser(in *pb.CreateOrUpdateUserReq) (*pb.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.BaseResp{}, nil
}
