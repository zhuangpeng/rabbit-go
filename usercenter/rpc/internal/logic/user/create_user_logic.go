package userlogic

import (
	"context"
	"strings"

	"github.com/zhuangpeng/rabbit-go/pkg/i18n"
	"github.com/zhuangpeng/rabbit-go/pkg/utils"
	"github.com/zhuangpeng/rabbit-go/pkg/utils/uuidx"
	"github.com/zhuangpeng/rabbit-go/pkg/xerr/statuserr"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/model"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
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

func (l *CreateUserLogic) CreateUser(in *pb.CreateUserReq) (*pb.BaseResp, error) {
	user := model.User{
		Name    : in.Name,
		Password: utils.BcryptEncrypt(in.Password),
		Nickname: in.Nickname,
		Mobile  : in.Mobile,
		Email   : in.Email,
		Avatar  : in.Avatar,
	}

	var (
		err     error
		resperr []string
	)
	// username不可重复
	_, err = l.svcCtx.UserModel.FindOneByName(l.ctx, user.Name)
	if err != nil && err != model.ErrNotFound {
		logx.Errorw(err.Error(), logx.Field("query username has err:", user.Name))
		return nil, statuserr.NewInvalidArgumentError(i18n.DatabaseError)
	}
	if err != model.ErrNotFound {
		resperr = append(resperr, i18n.UserAlreadyExist)
	}

	// mobile不可重复
	_, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, user.Mobile)
	if err != nil && err != model.ErrNotFound {
		logx.Errorw(err.Error(), logx.Field("query mobile has err:", user.Mobile))
		return nil, statuserr.NewInvalidArgumentError(i18n.DatabaseError)
	}
	if err != model.ErrNotFound {
		resperr = append(resperr, i18n.MobileAlreadyExist)
	}

	// email不可重复
	_, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, user.Email)
	if err != nil && err != model.ErrNotFound {
		logx.Errorw(err.Error(), logx.Field("query email has err:", user.Email))
		return nil, statuserr.NewInvalidArgumentError(i18n.DatabaseError)
	}
	if err != model.ErrNotFound {
		resperr = append(resperr, i18n.EmailAlreadyExist)
	}

	// 合并数组中的错误信息并返回
	if len(resperr) > 0 {
		errstrings := strings.Join(stringx.Remove(resperr, ""), "|")
		return nil, statuserr.NewAlreadyExistsError(errstrings)
	}

	// 创建用户信息
	user.Id = uuidx.NewUUID().String()
	result, err := l.svcCtx.UserModel.Insert(l.ctx, true, nil, &user)
	if err != nil {
		logx.Errorw(err.Error(), logx.Field("insert user err", user))
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}
	logx.Infof("the result is: %+v", result)

	return &pb.BaseResp{
		Msg: i18n.Success,
	}, nil
}
