package userlogic

import (
	"context"
	"strings"

	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/zhuangpeng/rabbit-go/pkg/i18n"
	"github.com/zhuangpeng/rabbit-go/pkg/utils"
	"github.com/zhuangpeng/rabbit-go/pkg/utils/uuidx"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/model"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
)

type CreateOrUpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateUserLogic {
	return &CreateOrUpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateUserLogic) CreateOrUpdateUser(in *pb.CreateOrUpdateUserReq) (*pb.BaseResp, error) {

	user := model.User{
		Name:     in.Username,
		Password: utils.BcryptEncrypt(in.Password) ,
		Nickname: in.Nickname,
		RoleId:   in.RoleId,
		Mobile:   in.Mobile,
		Email:    in.Email,
	}



	// 如果id为空则为创建用户
	if in.Id == "" {
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
		result, err := l.svcCtx.UserModel.InsertWithoutZero(l.ctx, &user)
		if err != nil {
			logx.Errorw(err.Error(), logx.Field("insert user err", user))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
		logx.Infof("the result is: %+v", result)
	}

	return &pb.BaseResp{}, nil
}
