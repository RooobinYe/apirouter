package logic

import (
	"context"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"
	"apirouter/rpc/apikey/apikeyclient"

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
	// 从JWT context中获取用户ID
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return &types.DeleteApiKeyResp{
			Code:    401,
			Message: "未授权访问",
		}, nil
	}

	// 调用RPC服务删除API密钥
	rpcResp, err := l.svcCtx.ApiKeyClient.DeleteApiKey(l.ctx, &apikeyclient.DeleteApiKeyRequest{
		UserId: userId,
		Id:     req.Id,
	})
	if err != nil {
		l.Logger.Errorf("DeleteApiKey RPC call failed: %v", err)
		return &types.DeleteApiKeyResp{
			Code:    500,
			Message: "删除API密钥失败",
		}, err
	}

	// 处理RPC响应
	if rpcResp.Code != 200 {
		return &types.DeleteApiKeyResp{
			Code:    int(rpcResp.Code),
			Message: rpcResp.Message,
		}, nil
	}

	return &types.DeleteApiKeyResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
