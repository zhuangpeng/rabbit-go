package menulogic

import (
	"context"

	"backend/rabbit-go/usercenter/rpc/internal/svc"
	"backend/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuParamListByMenuIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuParamListByMenuIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuParamListByMenuIdLogic {
	return &GetMenuParamListByMenuIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuParamListByMenuIdLogic) GetMenuParamListByMenuId(in *pb.IDReq) (*pb.MenuParamListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.MenuParamListResp{}, nil
}
