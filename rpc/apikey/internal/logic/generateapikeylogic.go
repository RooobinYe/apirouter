package logic

import (
	"context"

	"apirouter/rpc/apikey/apikey"
	"apirouter/rpc/apikey/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateApiKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateApiKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateApiKeyLogic {
	return &GenerateApiKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成API密钥
func (l *GenerateApiKeyLogic) GenerateApiKey(in *apikey.GenerateApiKeyRequest) (*apikey.GenerateApiKeyResponse, error) {
	// todo: add your logic here and delete this line

	return &apikey.GenerateApiKeyResponse{}, nil
}
