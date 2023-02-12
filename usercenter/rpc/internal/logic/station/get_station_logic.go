package stationlogic

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStationLogic {
	return &GetStationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStationLogic) GetStation(in *pb.IDReq) (*pb.GetStationResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetStationResp{}, nil
}
