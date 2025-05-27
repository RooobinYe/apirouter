package logic

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"apirouter/rpc/apikey/apikey"
	"apirouter/rpc/apikey/internal/svc"
	"apirouter/rpc/model"

	"github.com/google/uuid"
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
	// 生成UUID作为密钥ID
	id := uuid.New().String()

	// 生成16字节的随机密钥
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		l.Errorf("Failed to generate random bytes: %v", err)
		return &apikey.GenerateApiKeyResponse{
			Code:    500,
			Message: "生成API密钥失败",
		}, nil
	}

	// 将随机字节转换为十六进制字符串，添加前缀
	generatedKey := "sk-proj-" + hex.EncodeToString(bytes)

	// 创建数据库记录
	apiKeyData := &model.Apikeys{
		Id:        id,
		UserId:    in.UserId,
		Name:      in.Name,
		ApiKey:    generatedKey,
		CreatedAt: time.Now().Unix(),
		Status:    "active",
	}

	// 插入数据库
	_, err = l.svcCtx.ApiKeysModel.Insert(l.ctx, apiKeyData)
	if err != nil {
		l.Errorf("Failed to insert api key: %v", err)
		return &apikey.GenerateApiKeyResponse{
			Code:    500,
			Message: "保存API密钥失败",
		}, nil
	}

	// 构造响应
	apiKeyInfo := &apikey.ApiKeyInfo{
		Id:          id,
		UserId:      in.UserId,
		Name:        in.Name,
		ApiKey:      generatedKey,
		CreatedAt:   apiKeyData.CreatedAt,
		Status:      "active",
		ExpiresAt:   in.ExpiresAt,
		Description: in.Description,
	}

	return &apikey.GenerateApiKeyResponse{
		Code:    200,
		Message: "API密钥生成成功",
		Data:    apiKeyInfo,
	}, nil
}
