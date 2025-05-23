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
	// todo: add your logic here and delete this line

	return &apikey.ListApiKeysResponse{}, nil
}
