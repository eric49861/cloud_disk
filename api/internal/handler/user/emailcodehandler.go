package user

import (
	"cloud-disk/api/internal/logic/user"
	"cloud-disk/api/internal/svc"
	"cloud-disk/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EmailCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewEmailCodeLogic(r.Context(), svcCtx)
		resp, err := l.EmailCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
