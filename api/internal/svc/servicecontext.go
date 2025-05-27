package svc

import (
	"apirouter/api/internal/config"
	"apirouter/api/internal/middleware"
	"apirouter/rpc/apikey/apikeyclient"
	"apirouter/rpc/openai/openaiclient"
	"apirouter/rpc/user/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	AuthMiddleware   rest.Middleware
	ApiKeyMiddleware rest.Middleware
	UserClient       userclient.User
	ApiKeyClient     apikeyclient.ApiKey
	OpenAIClient     openaiclient.OpenAI
}

func NewServiceContext(c config.Config) *ServiceContext {
	userClient := userclient.NewUser(zrpc.MustNewClient(c.User))
	apiKeyClient := apikeyclient.NewApiKey(zrpc.MustNewClient(c.Apikey))
	
	return &ServiceContext{
		Config:           c,
		AuthMiddleware:   middleware.NewAuthMiddleware(userClient).Handle,
		ApiKeyMiddleware: middleware.NewApiKeyMiddleware(apiKeyClient).Handle,
		UserClient:       userClient,
		ApiKeyClient:     apiKeyClient,
		OpenAIClient:     openaiclient.NewOpenAI(zrpc.MustNewClient(c.Openai)),
	}
}
