package userlogic

import (
	"context"

	globalkey "github.com/zhuangpeng/rabbit-go/pkg/globalKey"
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

	var (
		user *model.User
		err  error
	)

	switch in.AuthType {
	case globalkey.UserAuthTypeMobile:
		if user, err = l.loginByMobile(in.AuthKey, in.Password); err != nil {
			return nil, err
		}
	case globalkey.UserAuthTypeEmail:
		if user, err = l.loginByEmail(in.AuthKey, in.Password); err != nil {
			return nil, err
		}
	case globalkey.UserAuthIDCard:
		if user, err = l.loginByIdCard(in.AuthKey, in.Password); err != nil {
			return nil, err
		}
	default:
		logx.Errorw("error login auth type", logx.Field("AuthType:", in.AuthType))
		return nil, status.Error(codes.InvalidArgument, "login.authTypeErr")
	}

	if !utils.BcryptCheck(in.Password, user.Password) {
		logx.Infof("password error, the user is: %s", user.Mobile)
		return nil, status.Error(codes.PermissionDenied, i18n.WrongPassword)
	}

	return &pb.LoginResp{}, nil
}

// 手机号登录
func (l *LoginLogic) loginByMobile(mobile, password string) (*model.User, error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Infof("fail to find auth mobile", logx.Field("mobile:", mobile))
			return nil, status.Error(codes.NotFound, i18n.MobileAlreadyExist)
		}
		logx.Errorw("fail to query user mobile when login", logx.Field("detail:", err))
		return nil, status.Error(codes.Internal, i18n.DatabaseError)
	}
	return user, nil
}

// 邮箱登录
func (l *LoginLogic) loginByEmail(email, password string) (*model.User, error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Infof("fail to find auth email", logx.Field("email:", email))
			return nil, status.Error(codes.NotFound, i18n.EmailAlreadyExist)
		}
		logx.Errorw("fail to query user email when login", logx.Field("detail:", err))
		return nil, status.Error(codes.Internal, i18n.DatabaseError)
	}
	return user, nil
}

// 身份证号登录
func (l *LoginLogic) loginByIdCard(idCard, password string) (*model.User, error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, idCard)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Infof("fail to find auth idcard", logx.Field("idcard:", idCard))
			return nil, status.Error(codes.NotFound, i18n.IdCardAlreadyExist)
		}
		logx.Errorw("fail to query user idcard when login", logx.Field("detail:", err))
		return nil, status.Error(codes.Internal, i18n.DatabaseError)
	}
	return user, nil
}