package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhuangpeng/rabbit-go/usercenter/api/internal/logic/user"
	"github.com/zhuangpeng/rabbit-go/usercenter/api/internal/svc"
	"github.com/zhuangpeng/rabbit-go/usercenter/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewRegisterLogic(r, svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
