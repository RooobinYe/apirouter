package logic

import (
	"context"

	"apirouter/rpc/apikey/apikey"
	"apirouter/rpc/apikey/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateKeyLogic {
	return &ValidateKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证API密钥（供其他服务调用）
func (l *ValidateKeyLogic) ValidateKey(in *apikey.ValidateKeyRequest) (*apikey.ValidateKeyResponse, error) {
	// todo: add your logic here and delete this line

	return &apikey.ValidateKeyResponse{}, nil
}
