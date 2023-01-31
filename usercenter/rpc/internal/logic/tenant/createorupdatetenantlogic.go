package tenantlogic

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateTenantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateTenantLogic {
	return &CreateOrUpdateTenantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateTenantLogic) CreateOrUpdateTenant(in *pb.CreateOrUpdateTenantReq) (*pb.IDResp, error) {
	// todo: add your logic here and delete this line

	return &pb.IDResp{}, nil
}
