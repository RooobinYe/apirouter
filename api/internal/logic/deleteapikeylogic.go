package logic

import (
	"context"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除API密钥
func NewDeleteApiKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiKeyLogic {
	return &DeleteApiKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApiKeyLogic) DeleteApiKey(req *types.DeleteApiKeyReq) (resp *types.DeleteApiKeyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
