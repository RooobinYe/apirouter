package logic

import (
	"context"

	"apirouter/rpc/user/internal/svc"
	"apirouter/rpc/user/user"

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
	// todo: add your logic here and delete this line

	return &user.ValidateTokenResponse{}, nil
}
