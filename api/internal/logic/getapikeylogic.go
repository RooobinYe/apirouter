package logic

import (
	"context"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"
	"apirouter/rpc/apikey/apikeyclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取API密钥详情
func NewGetApiKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiKeyLogic {
	return &GetApiKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiKeyLogic) GetApiKey(req *types.GetApiKeyReq) (resp *types.GetApiKeyResp, err error) {
	// 从JWT context中获取用户ID
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return &types.GetApiKeyResp{
			Code:    401,
			Message: "未授权访问",
		}, nil
	}

	// 调用RPC服务获取API密钥详情
	rpcResp, err := l.svcCtx.ApiKeyClient.GetApiKey(l.ctx, &apikeyclient.GetApiKeyRequest{
		UserId: userId,
		Id:     req.Id,
	})
	if err != nil {
		l.Logger.Errorf("GetApiKey RPC call failed: %v", err)
		return &types.GetApiKeyResp{
			Code:    500,
			Message: "获取API密钥失败",
		}, err
	}

	// 处理RPC响应
	if rpcResp.Code != 200 {
		return &types.GetApiKeyResp{
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

	return &types.GetApiKeyResp{
		Code:    200,
		Message: "获取成功",
		Data:    apiKeyInfo,
	}, nil
}
