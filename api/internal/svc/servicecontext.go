package svc

import (
	"apirouter/api/internal/config"
	"apirouter/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	AuthMiddleware   rest.Middleware
	ApiKeyMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		AuthMiddleware:   middleware.NewAuthMiddleware().Handle,
		ApiKeyMiddleware: middleware.NewApiKeyMiddleware().Handle,
	}
}
