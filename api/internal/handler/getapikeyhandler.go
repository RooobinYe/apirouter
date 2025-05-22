package handler

import (
	"net/http"

	"apirouter/api/internal/logic"
	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取API密钥详情
func GetApiKeyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetApiKeyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetApiKeyLogic(r.Context(), svcCtx)
		resp, err := l.GetApiKey(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
