package svc

import (
	"apirouter/rpc/apikey/apikeyclient"
	"apirouter/rpc/model"
	"apirouter/rpc/openai/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	ApiKeysModel model.ApikeysModel
	ApiKeyClient apikeyclient.ApiKey
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 启动时加载密钥配置
	if err := c.LoadSecrets(); err != nil {
		panic("Failed to load secrets: " + err.Error())
	}

	return &ServiceContext{
		Config:       c,
		ApiKeysModel: model.NewApikeysModel(sqlx.NewMysql(c.DataSource), c.Cache),
		ApiKeyClient: apikeyclient.NewApiKey(zrpc.MustNewClient(c.Apikey)),
	}
}
