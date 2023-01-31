package userlogic

import (
	"context"

	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/model"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	sq "github.com/Masterminds/squirrel"
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
	// 创建user对象初始化传入参数
	user := &model.User{
		Name    : in.Username,
		Password: in.Password,
		Nickname: in.Nickname,
		RoleId  : in.RoleId,
		Mobile  : in.Mobile,
		Email   : in.Email,
	}
	// 如果id为空则为创建用户
	if in.Id == "" {
		existUser := sq.Select("username,mobile,email").From("user").Where(sq.Eq{
			"username": user.Name,
		})
		logx.Infof("the query user is: %+v", existUser)
	}

	return &pb.BaseResp{}, nil
}
