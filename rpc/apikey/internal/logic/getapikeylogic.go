package logic

import (
	"context"

	"apirouter/rpc/apikey/apikey"
	"apirouter/rpc/apikey/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiKeyLogic {
	return &GetApiKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取API密钥详情
func (l *GetApiKeyLogic) GetApiKey(in *apikey.GetApiKeyRequest) (*apikey.GetApiKeyResponse, error) {
	// todo: add your logic here and delete this line

	return &apikey.GetApiKeyResponse{}, nil
}
