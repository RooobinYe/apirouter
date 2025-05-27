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
	// 查询API密钥
	apiKeyData, err := l.svcCtx.ApiKeysModel.FindOneByIdAndUserId(l.ctx, in.Id, in.UserId)
	if err != nil {
		l.Errorf("Failed to find api key: %v", err)
		return &apikey.GetApiKeyResponse{
			Code:    404,
			Message: "API密钥不存在",
		}, nil
	}

	// 为了安全，部分隐藏API密钥
	maskedApiKey := apiKeyData.ApiKey
	if len(apiKeyData.ApiKey) > 10 {
		maskedApiKey = apiKeyData.ApiKey[:10] + "..."
	}

	// 构造响应数据
	apiKeyInfo := &apikey.ApiKeyInfo{
		Id:        apiKeyData.Id,
		UserId:    apiKeyData.UserId,
		Name:      apiKeyData.Name,
		ApiKey:    maskedApiKey,
		CreatedAt: apiKeyData.CreatedAt,
		Status:    apiKeyData.Status,
	}

	return &apikey.GetApiKeyResponse{
		Code:    200,
		Message: "查询成功",
		Data:    apiKeyInfo,
	}, nil
}
