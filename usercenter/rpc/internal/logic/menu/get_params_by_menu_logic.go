package menulogic

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetParamsByMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetParamsByMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetParamsByMenuLogic {
	return &GetParamsByMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetParamsByMenuLogic) GetParamsByMenu(in *pb.IDReq) (*pb.GetParamsByMenuResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetParamsByMenuResp{}, nil
}
