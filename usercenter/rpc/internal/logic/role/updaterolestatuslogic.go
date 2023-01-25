package rolelogic

import (
	"context"

	"backend/rabbit-go/usercenter/rpc/internal/svc"
	"backend/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleStatusLogic {
	return &UpdateRoleStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleStatusLogic) UpdateRoleStatus(in *pb.StatusCodeReq) (*pb.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.BaseResp{}, nil
}
