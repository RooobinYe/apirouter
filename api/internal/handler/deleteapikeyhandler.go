package handler

import (
	"net/http"

	"apirouter/api/internal/logic"
	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除API密钥
func DeleteApiKeyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteApiKeyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteApiKeyLogic(r.Context(), svcCtx)
		resp, err := l.DeleteApiKey(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
