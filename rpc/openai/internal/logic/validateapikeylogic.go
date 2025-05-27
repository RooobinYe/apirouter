package logic

import (
	"context"

	"apirouter/rpc/apikey/apikeyclient"
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
	// 调用apikey RPC服务验证API密钥
	rpcResp, err := l.svcCtx.ApiKeyClient.ValidateKey(l.ctx, &apikeyclient.ValidateKeyRequest{
		ApiKey: in.ApiKey,
	})
	if err != nil {
		l.Errorf("Failed to validate api key: %v", err)
		return &openai.ValidateApiKeyResponse{
			Code:    500,
			Message: "验证API密钥失败",
		}, nil
	}

	// 处理RPC响应
	if rpcResp.Code != 200 {
		return &openai.ValidateApiKeyResponse{
			Code:    int32(rpcResp.Code),
			Message: rpcResp.Message,
		}, nil
	}

	// 构造响应数据
	validationData := &openai.ApiKeyValidationData{
		Valid:  rpcResp.Data.Valid,
		KeyId:  rpcResp.Data.KeyId,
		UserId: rpcResp.Data.UserId,
		Status: rpcResp.Data.Status,
	}

	return &openai.ValidateApiKeyResponse{
		Code:    200,
		Message: "验证完成",
		Data:    validationData,
	}, nil
}
