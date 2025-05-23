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
	// todo: add your logic here and delete this line

	return &apikey.DeleteApiKeyResponse{}, nil
}
