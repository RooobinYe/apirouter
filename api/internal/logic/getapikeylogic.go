package logic

import (
	"context"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
