package logic

import (
	"context"

	"apirouter/rpc/apikey/apikey"
	"apirouter/rpc/apikey/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListApiKeysLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListApiKeysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListApiKeysLogic {
	return &ListApiKeysLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取API密钥列表
func (l *ListApiKeysLogic) ListApiKeys(in *apikey.ListApiKeysRequest) (*apikey.ListApiKeysResponse, error) {
	// 设置默认分页参数
	page := in.Page
	pageSize := in.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询API密钥总数
	total, err := l.svcCtx.ApiKeysModel.CountByUserId(l.ctx, in.UserId)
	if err != nil {
		l.Errorf("Failed to count api keys: %v", err)
		return &apikey.ListApiKeysResponse{
			Code:    500,
			Message: "查询API密钥总数失败",
		}, nil
	}

	// 如果没有记录，直接返回空列表
	if total == 0 {
		return &apikey.ListApiKeysResponse{
			Code:    200,
			Message: "查询成功",
			Data: &apikey.ApiKeyListData{
				ApiKeys: []*apikey.ApiKeyInfo{},
				Total:   total,
			},
		}, nil
	}

	// 查询API密钥列表
	apiKeysData, err := l.svcCtx.ApiKeysModel.ListByUserId(l.ctx, in.UserId, pageSize, offset)
	if err != nil {
		l.Errorf("Failed to query api keys: %v", err)
		return &apikey.ListApiKeysResponse{
			Code:    500,
			Message: "查询API密钥列表失败",
		}, nil
	}

	// 转换为响应格式
	var apiKeys []*apikey.ApiKeyInfo
	for _, item := range apiKeysData {
		// 为了安全，不返回完整的API密钥，只显示前缀
		maskedApiKey := item.ApiKey
		if len(item.ApiKey) > 10 {
			maskedApiKey = item.ApiKey[:10] + "..."
		}

		apiKeyInfo := &apikey.ApiKeyInfo{
			Id:        item.Id,
			UserId:    item.UserId,
			Name:      item.Name,
			ApiKey:    maskedApiKey,
			CreatedAt: item.CreatedAt,
			Status:    item.Status,
		}
		apiKeys = append(apiKeys, apiKeyInfo)
	}

	return &apikey.ListApiKeysResponse{
		Code:    200,
		Message: "查询成功",
		Data: &apikey.ApiKeyListData{
			ApiKeys: apiKeys,
			Total:   total,
		},
	}, nil
}
