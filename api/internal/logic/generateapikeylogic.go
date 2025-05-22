package logic

import (
	"context"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
