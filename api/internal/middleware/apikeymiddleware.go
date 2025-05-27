package middleware

import (
	"context"
	"net/http"
	"strings"

	"apirouter/rpc/apikey/apikeyclient"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type ApiKeyMiddleware struct {
	apiKeyClient apikeyclient.ApiKey
}

func NewApiKeyMiddleware(apiKeyClient apikeyclient.ApiKey) *ApiKeyMiddleware {
	return &ApiKeyMiddleware{
		apiKeyClient: apiKeyClient,
	}
}

func (m *ApiKeyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取 Authorization 头
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			httpx.WriteJson(w, http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "缺少 Authorization 头",
			})
			return
		}

		// 检查是否以 "Bearer " 开头
		if !strings.HasPrefix(authorization, "Bearer ") {
			httpx.WriteJson(w, http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "无效的 Authorization 格式",
			})
			return
		}

		// 提取 API 密钥
		apiKey := strings.TrimPrefix(authorization, "Bearer ")
		if apiKey == "" {
			httpx.WriteJson(w, http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "API 密钥不能为空",
			})
			return
		}

		// 验证 API 密钥
		resp, err := m.apiKeyClient.ValidateKey(r.Context(), &apikeyclient.ValidateKeyRequest{
			ApiKey: apiKey,
		})
		if err != nil {
			logx.Errorf("API 密钥验证失败: %v", err)
			httpx.WriteJson(w, http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "API 密钥验证失败",
			})
			return
		}

		if resp.Code != 200 {
			httpx.WriteJson(w, http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": resp.Message,
			})
			return
		}

		// 将验证通过的 API 密钥存入 context
		ctx := context.WithValue(r.Context(), "api_key", apiKey)
		ctx = context.WithValue(ctx, "user_id", resp.Data.UserId)

		// 传递给下一个处理器
		next(w, r.WithContext(ctx))
	}
}
