package stationlogic

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/pkg/i18n"
	"github.com/zhuangpeng/rabbit-go/pkg/statuserr"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/model"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStationLogic {
	return &CreateStationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStationLogic) CreateStation(in *pb.CreateStationReq) (*pb.BaseResp, error) {
	station := &model.Station{
		Name:     in.Name,
		TenantId: in.TenantId,
	}

	_, err := l.svcCtx.StationModel.Insert(l.ctx, false, nil, station)
	if err != nil {
		logx.Errorw(err.Error(), logx.Field("insert station has error: %+v", in))
		return nil, statuserr.NewInternalError(i18n.CreateFailed)
	}

	return &pb.BaseResp{Msg: i18n.Success}, nil
}
