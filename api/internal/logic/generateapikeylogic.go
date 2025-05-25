package logic

import (
	"context"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"
	"apirouter/rpc/apikey/apikeyclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateApiKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成单个API密钥
func NewGenerateApiKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateApiKeyLogic {
	return &GenerateApiKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateApiKeyLogic) GenerateApiKey(req *types.GenerateApiKeyReq) (resp *types.GenerateApiKeyResp, err error) {
	// 从JWT context中获取用户ID
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return &types.GenerateApiKeyResp{
			Code:    401,
			Message: "未授权访问",
		}, nil
	}

	// 调用RPC服务生成API密钥
	rpcResp, err := l.svcCtx.ApiKeyClient.GenerateApiKey(l.ctx, &apikeyclient.GenerateApiKeyRequest{
		UserId:      userId,
		Name:        req.Name,
		Description: req.Description,
		ExpiresAt:   req.ExpiresAt,
	})
	if err != nil {
		l.Logger.Errorf("GenerateApiKey RPC call failed: %v", err)
		return &types.GenerateApiKeyResp{
			Code:    500,
			Message: "生成API密钥失败",
		}, err
	}

	// 处理RPC响应
	if rpcResp.Code != 200 {
		return &types.GenerateApiKeyResp{
			Code:    int(rpcResp.Code),
			Message: rpcResp.Message,
		}, nil
	}

	// 转换响应数据
	apiKeyInfo := types.ApiKeyInfo{
		Id:          rpcResp.Data.Id,
		UserId:      rpcResp.Data.UserId,
		Name:        rpcResp.Data.Name,
		ApiKey:      rpcResp.Data.ApiKey,
		CreatedAt:   rpcResp.Data.CreatedAt,
		Status:      rpcResp.Data.Status,
	}

	return &types.GenerateApiKeyResp{
		Code:    200,
		Message: "API密钥生成成功",
		Data:    apiKeyInfo,
	}, nil
}
