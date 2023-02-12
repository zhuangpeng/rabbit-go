package stationlogic

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStationByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStationByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStationByUserLogic {
	return &GetStationByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStationByUserLogic) GetStationByUser(in *pb.IDReq) (*pb.GetStationByUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetStationByUserResp{}, nil
}
