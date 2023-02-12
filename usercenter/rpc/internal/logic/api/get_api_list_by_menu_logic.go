package apilogic

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListByMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiListByMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListByMenuLogic {
	return &GetApiListByMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApiListByMenuLogic) GetApiListByMenu(in *pb.IDReq) (*pb.GetApiListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetApiListResp{}, nil
}
