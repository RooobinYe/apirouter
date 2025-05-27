package logic

import (
	"context"

	"apirouter/rpc/user/internal/svc"
	"apirouter/rpc/user/user"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证Token（供AuthMiddleware使用）
func (l *ValidateTokenLogic) ValidateToken(in *user.ValidateTokenRequest) (*user.ValidateTokenResponse, error) {
	// 参数验证
	if in.AccessToken == "" {
		return &user.ValidateTokenResponse{
			Code:    400,
			Message: "访问令牌不能为空",
			Data: &user.TokenData{
				Valid: false,
			},
		}, nil
	}

	// 解析JWT Token
	token, err := jwt.Parse(in.AccessToken, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(l.svcCtx.Config.JwtAuth.AccessSecret), nil
	})

	if err != nil {
		l.Infof("Failed to parse token: %v", err)
		return &user.ValidateTokenResponse{
			Code:    200,
			Message: "验证完成",
			Data: &user.TokenData{
				Valid: false,
			},
		}, nil
	}

	// 检查token是否有效
	if !token.Valid {
		return &user.ValidateTokenResponse{
			Code:    200,
			Message: "验证完成",
			Data: &user.TokenData{
				Valid: false,
			},
		}, nil
	}

	// 获取claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		l.Error("Failed to get claims from token")
		return &user.ValidateTokenResponse{
			Code:    200,
			Message: "验证完成",
			Data: &user.TokenData{
				Valid: false,
			},
		}, nil
	}

	// 提取用户信息
	userIdFloat, ok := claims["user_id"].(float64)
	if !ok {
		l.Error("Failed to get user_id from claims")
		return &user.ValidateTokenResponse{
			Code:    200,
			Message: "验证完成",
			Data: &user.TokenData{
				Valid: false,
			},
		}, nil
	}

	username, ok := claims["username"].(string)
	if !ok {
		l.Error("Failed to get username from claims")
		return &user.ValidateTokenResponse{
			Code:    200,
			Message: "验证完成",
			Data: &user.TokenData{
				Valid: false,
			},
		}, nil
	}

	expiresAtFloat, ok := claims["exp"].(float64)
	if !ok {
		l.Error("Failed to get exp from claims")
		return &user.ValidateTokenResponse{
			Code:    200,
			Message: "验证完成",
			Data: &user.TokenData{
				Valid: false,
			},
		}, nil
	}

	// 构造验证结果
	tokenData := &user.TokenData{
		Valid:     true,
		UserId:    int64(userIdFloat),
		Username:  username,
		ExpiresAt: int64(expiresAtFloat),
	}

	return &user.ValidateTokenResponse{
		Code:    200,
		Message: "验证完成",
		Data:    tokenData,
	}, nil
}
