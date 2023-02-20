package tenantlogic

import (
	"context"
	"time"

	globalkey "github.com/zhuangpeng/rabbit-go/pkg/globalKey"
	"github.com/zhuangpeng/rabbit-go/pkg/i18n"
	"github.com/zhuangpeng/rabbit-go/pkg/statuserr"
	"github.com/zhuangpeng/rabbit-go/pkg/utils/dbx"
	"github.com/zhuangpeng/rabbit-go/pkg/utils/uuidx"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/model"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTenantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTenantLogic {
	return &CreateTenantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTenantLogic) CreateTenant(in *pb.CreateTenantReq) (*pb.BaseResp, error) {
	tenant := &model.Tenant{
		Id       : uuidx.NewUUID().String(),
		ParentId : in.ParentId,
		Name     : in.Name,
		Contact  : in.Contact,
		Mobile   : in.Mobile,
		Sort     : int64(in.SortNo),
		StartTime: time.UnixMilli(in.StartTime),
		Account  : in.AccountId,
	}

	if in.EndTime == 0 {
		tenant.EndTime = globalkey.MaxEndTime
	} else {
		tenant.EndTime = dbx.NewNullTime(in.EndTime)
	}

	

	if _, err := l.svcCtx.TenantModel.Insert(l.ctx, true, nil, tenant); err != nil {
		logx.Errorw(err.Error(), logx.Field("insert tenant info has error: %+v", tenant))
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	return &pb.BaseResp{Msg: i18n.Success}, nil
}
