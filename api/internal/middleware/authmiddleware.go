package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"apirouter/rpc/user/user"
	"apirouter/rpc/user/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthMiddleware struct {
	userRpc userclient.User
}

func NewAuthMiddleware(userRpc userclient.User) *AuthMiddleware {
	return &AuthMiddleware{
		userRpc: userRpc,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从Header中获取Authorization token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeErrorResponse(w, 401, "未提供认证信息")
			return
		}

		// 检查Bearer token格式
		if !strings.HasPrefix(authHeader, "Bearer ") {
			writeErrorResponse(w, 401, "认证格式错误")
			return
		}

		// 提取token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			writeErrorResponse(w, 401, "令牌不能为空")
			return
		}

		// 调用user RPC服务验证token
		resp, err := m.userRpc.ValidateToken(r.Context(), &user.ValidateTokenRequest{
			AccessToken: token,
		})

		if err != nil {
			logx.Errorf("Failed to validate token: %v", err)
			writeErrorResponse(w, 500, "验证服务异常")
			return
		}

		// 检查验证结果
		if resp.Data == nil || !resp.Data.Valid {
			writeErrorResponse(w, 401, "令牌无效或已过期")
			return
		}

		// 将用户信息添加到请求头中，供后续handler使用
		r.Header.Set("X-User-Id", strconv.FormatInt(resp.Data.UserId, 10))
		r.Header.Set("X-Username", resp.Data.Username)

		// 将用户信息添加到请求上下文中
		ctx := context.WithValue(r.Context(), "user_id", resp.Data.UserId)
		ctx = context.WithValue(ctx, "username", resp.Data.Username)

		// 通过验证，继续执行下一个handler
		next(w, r.WithContext(ctx))
	}
}

func writeErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    nil,
	}

	json.NewEncoder(w).Encode(response)
}
