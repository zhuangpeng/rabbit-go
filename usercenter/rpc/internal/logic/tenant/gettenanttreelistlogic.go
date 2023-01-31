package tenantlogic

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTenantTreeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTenantTreeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTenantTreeListLogic {
	return &GetTenantTreeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTenantTreeListLogic) GetTenantTreeList(in *pb.TenantListReq) (*pb.TenantListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.TenantListResp{}, nil
}
