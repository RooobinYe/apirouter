package logic

import (
	"context"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenAIChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// OpenAI Chat 聊天接口
func NewOpenAIChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenAIChatLogic {
	return &OpenAIChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenAIChatLogic) OpenAIChat(req *types.OpenAIChatReq) (resp *types.OpenAIChatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
