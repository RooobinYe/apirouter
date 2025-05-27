package logic

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"time"

	"apirouter/rpc/apikey/apikeyclient"
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

// OpenAI API 请求结构
type OpenAIAPIRequest struct {
	Model            string             `json:"model"`
	Messages         []OpenAIAPIMessage `json:"messages"`
	Temperature      *float64           `json:"temperature,omitempty"`
	MaxTokens        *int32             `json:"max_tokens,omitempty"`
	TopP             *float64           `json:"top_p,omitempty"`
	FrequencyPenalty *float64           `json:"frequency_penalty,omitempty"`
	PresencePenalty  *float64           `json:"presence_penalty,omitempty"`
	Stop             []string           `json:"stop,omitempty"`
	Stream           bool               `json:"stream,omitempty"`
	User             string             `json:"user,omitempty"`
}

type OpenAIAPIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name,omitempty"`
}

// OpenAI API 响应结构
type OpenAIAPIResponse struct {
	ID      string            `json:"id"`
	Object  string            `json:"object"`
	Created int64             `json:"created"`
	Model   string            `json:"model"`
	Choices []OpenAIAPIChoice `json:"choices"`
	Usage   OpenAIAPIUsage    `json:"usage"`
}

type OpenAIAPIChoice struct {
	Index        int              `json:"index"`
	Message      OpenAIAPIMessage `json:"message"`
	FinishReason string           `json:"finish_reason"`
}

type OpenAIAPIUsage struct {
	PromptTokens     int32 `json:"prompt_tokens"`
	CompletionTokens int32 `json:"completion_tokens"`
	TotalTokens      int32 `json:"total_tokens"`
}

// 生成随机ID
func generateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return fmt.Sprintf("chatcmpl-%x", bytes[:8])
}

// OpenAI Chat 聊天接口
func (l *ChatCompletionsLogic) ChatCompletions(in *openai.OpenAIChatRequest) (*openai.OpenAIChatResponse, error) {
	// 1. 验证API密钥
	if in.ApiKey == "" {
		return &openai.OpenAIChatResponse{
			Code:    400,
			Message: "API密钥不能为空",
		}, nil
	}

	// 调用验证API密钥服务
	validateResp, err := l.svcCtx.ApiKeyClient.ValidateKey(l.ctx, &apikeyclient.ValidateKeyRequest{
		ApiKey: in.ApiKey,
	})
	if err != nil {
		l.Errorf("Failed to validate api key: %v", err)
		return &openai.OpenAIChatResponse{
			Code:    500,
			Message: "验证API密钥失败",
		}, nil
	}

	if validateResp.Code != 200 || !validateResp.Data.Valid {
		return &openai.OpenAIChatResponse{
			Code:    401,
			Message: "API密钥无效",
		}, nil
	}

	// 2. 直接转发用户的原始请求到OpenAI API
	// 从密钥文件获取OpenAI API密钥
	openaiAPIKey := l.svcCtx.Config.GetOpenAIAPIKey()
	if openaiAPIKey == "" {
		l.Error("OpenAI API密钥未配置，请在 etc/secrets.yaml 文件中配置有效的密钥")
		return &openai.OpenAIChatResponse{
			Code:    500,
			Message: "OpenAI API密钥未配置",
		}, nil
	}

	req, err := http.NewRequestWithContext(l.ctx, "POST", "https://api.openai.com/v1/chat/completions", bytes.NewBufferString(in.RequestBody))
	if err != nil {
		l.Errorf("Failed to create request: %v", err)
		return &openai.OpenAIChatResponse{
			Code:    500,
			Message: "创建请求失败",
		}, nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openaiAPIKey)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		l.Errorf("Failed to call OpenAI API: %v", err)
		return &openai.OpenAIChatResponse{
			Code:    500,
			Message: "调用OpenAI API失败",
		}, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		l.Errorf("Failed to read response body: %v", err)
		return &openai.OpenAIChatResponse{
			Code:    500,
			Message: "读取响应失败",
		}, nil
	}

	if resp.StatusCode != http.StatusOK {
		l.Errorf("OpenAI API returned error: %s", string(body))
		return &openai.OpenAIChatResponse{
			Code:    int32(resp.StatusCode),
			Message: "OpenAI API请求失败",
		}, nil
	}

	// 3. 直接返回OpenAI的原始响应（作为JSON字符串）
	return &openai.OpenAIChatResponse{
		Code:    200,
		Message: "请求成功",
		Data: &openai.OpenAIChatData{
			RawResponse: string(body), // 添加原始响应字段
		},
	}, nil
}
