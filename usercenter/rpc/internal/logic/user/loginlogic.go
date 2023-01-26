package userlogic

import (
	"context"

	"backend/rabbit-go/pkg/xerr"
	"backend/rabbit-go/usercenter/rpc/internal/model"
	"backend/rabbit-go/usercenter/rpc/internal/svc"
	"backend/rabbit-go/usercenter/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	logx.Infof("the user model object is: %+v", l.svcCtx.UserModel)
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	logx.Infof("the user is: %+v", user)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "用户不存在", in.Username, err)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "user not found", in.Username, err)
	}
	// userModel := new(model.User)
	// userModel.Username = user.Username
	return &pb.LoginResp{
		Id:        user.Username,
		RoleName:  "rolename",
		RoleValue: "rolevalue",
		RoleId:    0,
	}, nil
}
