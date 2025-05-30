// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3
// Source: apikey.proto

package server

import (
	"context"

	"apirouter/rpc/apikey/apikey"
	"apirouter/rpc/apikey/internal/logic"
	"apirouter/rpc/apikey/internal/svc"
)

type ApiKeyServer struct {
	svcCtx *svc.ServiceContext
	apikey.UnimplementedApiKeyServer
}

func NewApiKeyServer(svcCtx *svc.ServiceContext) *ApiKeyServer {
	return &ApiKeyServer{
		svcCtx: svcCtx,
	}
}

// 生成API密钥
func (s *ApiKeyServer) GenerateApiKey(ctx context.Context, in *apikey.GenerateApiKeyRequest) (*apikey.GenerateApiKeyResponse, error) {
	l := logic.NewGenerateApiKeyLogic(ctx, s.svcCtx)
	return l.GenerateApiKey(in)
}

// 获取API密钥列表
func (s *ApiKeyServer) ListApiKeys(ctx context.Context, in *apikey.ListApiKeysRequest) (*apikey.ListApiKeysResponse, error) {
	l := logic.NewListApiKeysLogic(ctx, s.svcCtx)
	return l.ListApiKeys(in)
}

// 获取API密钥详情
func (s *ApiKeyServer) GetApiKey(ctx context.Context, in *apikey.GetApiKeyRequest) (*apikey.GetApiKeyResponse, error) {
	l := logic.NewGetApiKeyLogic(ctx, s.svcCtx)
	return l.GetApiKey(in)
}

// 删除API密钥
func (s *ApiKeyServer) DeleteApiKey(ctx context.Context, in *apikey.DeleteApiKeyRequest) (*apikey.DeleteApiKeyResponse, error) {
	l := logic.NewDeleteApiKeyLogic(ctx, s.svcCtx)
	return l.DeleteApiKey(in)
}

// 验证API密钥（供其他服务调用）
func (s *ApiKeyServer) ValidateKey(ctx context.Context, in *apikey.ValidateKeyRequest) (*apikey.ValidateKeyResponse, error) {
	l := logic.NewValidateKeyLogic(ctx, s.svcCtx)
	return l.ValidateKey(in)
}

// 更新API密钥状态（内部服务使用）
func (s *ApiKeyServer) UpdateKeyStatus(ctx context.Context, in *apikey.UpdateKeyStatusRequest) (*apikey.UpdateKeyStatusResponse, error) {
	l := logic.NewUpdateKeyStatusLogic(ctx, s.svcCtx)
	return l.UpdateKeyStatus(in)
}
