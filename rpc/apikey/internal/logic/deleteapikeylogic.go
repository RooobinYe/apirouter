package logic

import (
	"context"

	"apirouter/rpc/apikey/apikey"
	"apirouter/rpc/apikey/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteApiKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiKeyLogic {
	return &DeleteApiKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除API密钥
func (l *DeleteApiKeyLogic) DeleteApiKey(in *apikey.DeleteApiKeyRequest) (*apikey.DeleteApiKeyResponse, error) {
	// 检查API密钥是否存在且属于当前用户
	_, err := l.svcCtx.ApiKeysModel.FindOneByIdAndUserId(l.ctx, in.Id, in.UserId)
	if err != nil {
		l.Errorf("Failed to find api key: %v", err)
		return &apikey.DeleteApiKeyResponse{
			Code:    404,
			Message: "API密钥不存在",
		}, nil
	}

	// 删除API密钥
	err = l.svcCtx.ApiKeysModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Errorf("Failed to delete api key: %v", err)
		return &apikey.DeleteApiKeyResponse{
			Code:    500,
			Message: "删除API密钥失败",
		}, nil
	}

	return &apikey.DeleteApiKeyResponse{
		Code:    200,
		Message: "API密钥删除成功",
		Data: &apikey.SuccessData{
			Message: "删除成功",
		},
	}, nil
}
