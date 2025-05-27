package handler

import (
	"io"
	"net/http"

	"apirouter/api/internal/logic"
	"apirouter/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// OpenAI Chat 聊天接口
func OpenAIChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 读取原始请求体
		body, err := io.ReadAll(r.Body)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOpenAIChatLogic(r.Context(), svcCtx)
		resp, err := l.OpenAIChat(body)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
