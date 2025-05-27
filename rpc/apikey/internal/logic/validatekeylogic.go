package logic

import (
	"context"

	"apirouter/rpc/apikey/apikey"
	"apirouter/rpc/apikey/internal/svc"
	"apirouter/rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateKeyLogic {
	return &ValidateKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证API密钥（供其他服务调用）
func (l *ValidateKeyLogic) ValidateKey(in *apikey.ValidateKeyRequest) (*apikey.ValidateKeyResponse, error) {
	// 查询API密钥
	apiKeyData, err := l.svcCtx.ApiKeysModel.FindOneByApiKey(l.ctx, in.ApiKey)
	if err != nil {
		if err == model.ErrNotFound {
			// 密钥不存在
			return &apikey.ValidateKeyResponse{
				Code:    200,
				Message: "验证完成",
				Data: &apikey.KeyValidationData{
					Valid: false,
				},
			}, nil
		}
		
		l.Errorf("Failed to find api key: %v", err)
		return &apikey.ValidateKeyResponse{
			Code:    500,
			Message: "验证API密钥失败",
		}, nil
	}

	// 检查密钥状态
	valid := apiKeyData.Status == "active"

	// 构造验证结果
	validationData := &apikey.KeyValidationData{
		Valid:  valid,
		KeyId:  apiKeyData.Id,
		UserId: apiKeyData.UserId,
		Status: apiKeyData.Status,
		Name:   apiKeyData.Name,
	}

	return &apikey.ValidateKeyResponse{
		Code:    200,
		Message: "验证完成",
		Data:    validationData,
	}, nil
}
