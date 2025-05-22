package logic

import (
	"context"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
