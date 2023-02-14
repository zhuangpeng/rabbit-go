package userlogic

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/pkg/i18n"
	"github.com/zhuangpeng/rabbit-go/pkg/utils"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/model"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	switch in.AuthType {
	case model.UserAuthTypeUserName:
	case model.UserAuthTypeMobile:
	default:
		logx.Errorw("error login auth type", logx.Field("AuthType:", in.AuthType))
		return nil, status.Error(codes.InvalidArgument, "login.authTypeErr")
	}

	return &pb.LoginResp{}, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (string, error) {

	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil {
		if err != model.ErrNotFound {
			logx.Errorw("fail to find auth mobile", logx.Field("mobile:", mobile))
			return "", status.Error(codes.NotFound, "login.mobileNotExist")
		}
		logx.Errorw("fail to query user when login", logx.Field("detail:", err))
		return "", status.Error(codes.Internal, i18n.DatabaseError)
	}

	if !utils.BcryptCheck(password, user.Password) {
		logx.Infof("password error, the user is: %s", user.Name)
	}

	return user.Id, nil
}

func (l *LoginLogic) loginBySmallWx() error {
	return nil
}
