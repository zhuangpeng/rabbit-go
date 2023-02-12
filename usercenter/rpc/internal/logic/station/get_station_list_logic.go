package stationlogic

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStationListLogic {
	return &GetStationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStationListLogic) GetStationList(in *pb.PageInfoReq) (*pb.GetStationListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetStationListResp{}, nil
}
