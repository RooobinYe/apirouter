package logic

import (
	"context"
	"errors"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"
	"apirouter/rpc/openai/openaiclient"

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

func (l *OpenAIChatLogic) OpenAIChat(req *types.OpenAIChatReq) (resp interface{}, err error) {
	// 从API Key中间件获取验证通过的API密钥
	apiKey, ok := l.ctx.Value("api_key").(string)
	if !ok {
		return nil, errors.New("API密钥验证失败")
	}

	// 转换请求消息格式
	var messages []*openaiclient.OpenAIChatMessage
	for _, msg := range req.Messages {
		messages = append(messages, &openaiclient.OpenAIChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
			Name:    msg.Name,
		})
	}

	// 调用OpenAI RPC服务
	rpcResp, err := l.svcCtx.OpenAIClient.ChatCompletions(l.ctx, &openaiclient.OpenAIChatRequest{
		Model:            req.Model,
		Messages:         messages,
		Temperature:      req.Temperature,
		MaxTokens:        int32(req.MaxTokens),
		TopP:             req.TopP,
		FrequencyPenalty: req.FrequencyPenalty,
		PresencePenalty:  req.PresencePenalty,
		Stop:             req.Stop,
		Stream:           req.Stream,
		User:             req.User,
		ApiKey:           apiKey,
	})
	if err != nil {
		l.Logger.Errorf("OpenAI ChatCompletions RPC call failed: %v", err)
		return nil, err
	}

	// 处理RPC响应
	if rpcResp.Code != 200 {
		return nil, errors.New(rpcResp.Message)
	}

	// 转换响应数据
	var choices []types.OpenAIChatChoice
	for _, choice := range rpcResp.Data.Choices {
		choices = append(choices, types.OpenAIChatChoice{
			Index: int(choice.Index),
			Message: types.OpenAIChatMessage{
				Role:    choice.Message.Role,
				Content: choice.Message.Content,
				Name:    choice.Message.Name,
			},
			FinishReason: choice.FinishReason,
		})
	}

	return &types.OpenAIChatResp{
		Id:      rpcResp.Data.Id,
		Object:  rpcResp.Data.Object,
		Created: rpcResp.Data.Created,
		Model:   rpcResp.Data.Model,
		Choices: choices,
		Usage: types.OpenAIUsage{
			PromptTokens:     int(rpcResp.Data.Usage.PromptTokens),
			CompletionTokens: int(rpcResp.Data.Usage.CompletionTokens),
			TotalTokens:      int(rpcResp.Data.Usage.TotalTokens),
		},
	}, nil
}
