package userlogic

import (
	"context"

	globalkey "github.com/zhuangpeng/rabbit-go/pkg/globalKey"
	"github.com/zhuangpeng/rabbit-go/pkg/i18n"
	"github.com/zhuangpeng/rabbit-go/pkg/statuserr"
	"github.com/zhuangpeng/rabbit-go/pkg/utils"
	"github.com/zhuangpeng/rabbit-go/pkg/utils/uuidx"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/model"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TODO: 小程序注册
func (l *CreateUserLogic) CreateUser(in *pb.CreateUserReq) (*pb.BaseResp, error) {

	var (
		user *model.User
		err  error
	)

	switch in.AuthType {
	case globalkey.UserAuthTypeMobile:
		if user, err = l.registerByMobile(in); err != nil {
			return nil, err
		}
	case globalkey.UserAuthTypeEmail:
		if user, err = l.registerByEmail(in); err != nil {
			return nil, err
		}
	case globalkey.UserAuthIDCard:
		if user, err = l.registerByIdCard(in); err != nil {
			return nil, err
		}
	default:
		logx.Errorw("error register auth type", logx.Field("AuthType:", in.AuthType))
		return nil, statuserr.NewInvalidArgumentError(i18n.ParamsError)
	}
	// 创建用户
	if _, err := l.svcCtx.UserModel.Insert(l.ctx, true, nil, user); err != nil {
		logx.Errorw(err.Error(), logx.Field("insert user register by idcard error: %+v", user))
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	return &pb.BaseResp{
		Msg: i18n.Success,
	}, nil
}

func (l *CreateUserLogic) registerByMobile(in *pb.CreateUserReq) (*model.User, error) {
	// 手机号不可重复注册
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.AuthKey)
	if err != nil {
		if err != model.ErrNotFound {
			logx.Infof(err.Error(), logx.Field("register mobile exist:", in.AuthKey))
			return nil, statuserr.NewAlreadyExistsError(i18n.MobileAlreadyExist)
		}
		logx.Errorw(err.Error(), logx.Field("register mobile as error:", in.AuthKey))
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}
	// 创建用户
	return &model.User{
		Id:       uuidx.NewUUID().String(),
		Mobile:   in.AuthKey,
		Password: utils.BcryptEncrypt(in.Password),
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
	}, nil
}

func (l *CreateUserLogic) registerByEmail(in *pb.CreateUserReq) (*model.User, error) {
	// 邮箱不可重复注册
	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.AuthKey)
	if err != nil {
		if err != model.ErrNotFound {
			logx.Infof(err.Error(), logx.Field("register email exist:", in.AuthKey))
			return nil, statuserr.NewAlreadyExistsError(i18n.MobileAlreadyExist)
		}
		logx.Errorw(err.Error(), logx.Field("register email as error:", in.AuthKey))
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}
	// 创建用户
	return &model.User{
		Id:       uuidx.NewUUID().String(),
		Email:    in.AuthKey,
		Password: utils.BcryptEncrypt(in.Password),
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
	}, nil
}

func (l *CreateUserLogic) registerByIdCard(in *pb.CreateUserReq) (*model.User, error) {
	// 身份证号不可重复注册
	_, err := l.svcCtx.UserModel.FindOneByIdCard(l.ctx, in.AuthKey)
	if err != nil {
		if err != model.ErrNotFound {
			logx.Infof(err.Error(), logx.Field("register idcard exist:", in.AuthKey))
			return nil, statuserr.NewAlreadyExistsError(i18n.MobileAlreadyExist)
		}
		logx.Errorw(err.Error(), logx.Field("register idcard as error:", in.AuthKey))
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}
	// 创建用户
	return &model.User{
		Id:       uuidx.NewUUID().String(),
		IdCard:   in.AuthKey,
		Password: utils.BcryptEncrypt(in.Password),
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
	}, nil
}
