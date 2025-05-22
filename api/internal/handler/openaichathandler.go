package handler

import (
	"net/http"

	"apirouter/api/internal/logic"
	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// OpenAI Chat 聊天接口
func OpenAIChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OpenAIChatReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOpenAIChatLogic(r.Context(), svcCtx)
		resp, err := l.OpenAIChat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
