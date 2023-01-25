package rolelogic

import (
	"context"

	"backend/rabbit-go/usercenter/rpc/internal/svc"
	"backend/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateRoleLogic {
	return &CreateOrUpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateRoleLogic) CreateOrUpdateRole(in *pb.RoleInfo) (*pb.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.BaseResp{}, nil
}
