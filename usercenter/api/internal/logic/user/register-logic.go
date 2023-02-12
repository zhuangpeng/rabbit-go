package user

import (
	"context"
	"net/http"

	"github.com/zhuangpeng/rabbit-go/usercenter/api/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/api/internal/types"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewRegisterLogic(r *http.Request, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.UserRpc.CreateUser(l.ctx, &pb.CreateUserReq{
		Name:     req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
		Mobile:   req.Mobile,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, data.Msg)}, nil
}
