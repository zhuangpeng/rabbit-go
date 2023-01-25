package menulogic

import (
	"context"

	"backend/rabbit-go/usercenter/rpc/internal/svc"
	"backend/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByRoleLogic {
	return &GetMenuListByRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuListByRoleLogic) GetMenuListByRole(in *pb.IDReq) (*pb.MenuInfoList, error) {
	// todo: add your logic here and delete this line

	return &pb.MenuInfoList{}, nil
}
