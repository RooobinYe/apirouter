package svc

import (
	"apirouter/rpc/apikey/internal/config"
	"apirouter/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ApiKeysModel model.ApikeysModel
	UsersModel   model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ApiKeysModel: model.NewApikeysModel(sqlx.NewMysql(c.DataSource), c.Cache),
		UsersModel:   model.NewUsersModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
