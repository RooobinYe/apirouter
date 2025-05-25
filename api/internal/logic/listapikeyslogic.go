package logic

import (
	"context"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"
	"apirouter/rpc/apikey/apikeyclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListApiKeysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取API密钥列表
func NewListApiKeysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListApiKeysLogic {
	return &ListApiKeysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListApiKeysLogic) ListApiKeys(req *types.ListApiKeysReq) (resp *types.ListApiKeysResp, err error) {
	// 从JWT context中获取用户ID
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return &types.ListApiKeysResp{
			Code:    401,
			Message: "未授权访问",
		}, nil
	}

	// 调用RPC服务获取API密钥列表
	rpcResp, err := l.svcCtx.ApiKeyClient.ListApiKeys(l.ctx, &apikeyclient.ListApiKeysRequest{
		UserId:   userId,
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	})
	if err != nil {
		l.Logger.Errorf("ListApiKeys RPC call failed: %v", err)
		return &types.ListApiKeysResp{
			Code:    500,
			Message: "获取API密钥列表失败",
		}, err
	}

	// 处理RPC响应
	if rpcResp.Code != 200 {
		return &types.ListApiKeysResp{
			Code:    int(rpcResp.Code),
			Message: rpcResp.Message,
		}, nil
	}

	// 转换响应数据
	var apiKeys []types.ApiKeyInfo
	for _, item := range rpcResp.Data.ApiKeys {
		apiKeys = append(apiKeys, types.ApiKeyInfo{
			Id:          item.Id,
			UserId:      item.UserId,
			Name:        item.Name,
			ApiKey:      item.ApiKey,
			CreatedAt:   item.CreatedAt,
			Status:      item.Status,
		})
	}

	return &types.ListApiKeysResp{
		Code:    200,
		Message: "获取成功",
		Data: types.ApiKeyListData{
			ApiKeys: apiKeys,
			Total:   rpcResp.Data.Total,
		},
	}, nil
}
