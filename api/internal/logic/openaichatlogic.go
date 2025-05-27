package logic

import (
	"context"
	"encoding/json"
	"errors"

	"apirouter/api/internal/svc"
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

func (l *OpenAIChatLogic) OpenAIChat(requestBody []byte) (resp interface{}, err error) {
	// 从API Key中间件获取验证通过的API密钥
	apiKey, ok := l.ctx.Value("api_key").(string)
	if !ok {
		return nil, errors.New("API密钥验证失败")
	}

	// 调用OpenAI RPC服务，直接传递原始请求体
	rpcResp, err := l.svcCtx.OpenAIClient.ChatCompletions(l.ctx, &openaiclient.OpenAIChatRequest{
		RequestBody: string(requestBody),
		ApiKey:      apiKey,
	})
	if err != nil {
		l.Logger.Errorf("OpenAI ChatCompletions RPC call failed: %v", err)
		return nil, err
	}

	// 处理RPC响应
	if rpcResp.Code != 200 {
		return nil, errors.New(rpcResp.Message)
	}

	// 直接返回OpenAI的原始响应JSON
	var rawResponse interface{}
	if err := json.Unmarshal([]byte(rpcResp.Data.RawResponse), &rawResponse); err != nil {
		l.Logger.Errorf("Failed to unmarshal raw response: %v", err)
		return nil, err
	}
	return rawResponse, nil
}
