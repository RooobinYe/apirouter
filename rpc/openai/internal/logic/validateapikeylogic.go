package logic

import (
	"context"

	"apirouter/rpc/openai/internal/svc"
	"apirouter/rpc/openai/openai"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateApiKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateApiKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateApiKeyLogic {
	return &ValidateApiKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证API密钥（供ApiKeyMiddleware使用）
func (l *ValidateApiKeyLogic) ValidateApiKey(in *openai.ValidateApiKeyRequest) (*openai.ValidateApiKeyResponse, error) {
	// todo: add your logic here and delete this line

	return &openai.ValidateApiKeyResponse{}, nil
}
