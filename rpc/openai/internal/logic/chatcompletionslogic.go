package logic

import (
	"context"

	"apirouter/rpc/openai/internal/svc"
	"apirouter/rpc/openai/openai"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatCompletionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatCompletionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatCompletionsLogic {
	return &ChatCompletionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OpenAI Chat 聊天接口
func (l *ChatCompletionsLogic) ChatCompletions(in *openai.OpenAIChatRequest) (*openai.OpenAIChatResponse, error) {
	// todo: add your logic here and delete this line

	return &openai.OpenAIChatResponse{}, nil
}
