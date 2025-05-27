package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ApikeysModel = (*customApikeysModel)(nil)

type (
	// ApikeysModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApikeysModel.
	ApikeysModel interface {
		apikeysModel
		CountByUserId(ctx context.Context, userId int64) (int64, error)
		ListByUserId(ctx context.Context, userId int64, limit, offset int32) ([]*Apikeys, error)
		FindOneByIdAndUserId(ctx context.Context, id string, userId int64) (*Apikeys, error)
	}

	customApikeysModel struct {
		*defaultApikeysModel
	}
)

// NewApikeysModel returns a model for the database table.
func NewApikeysModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ApikeysModel {
	return &customApikeysModel{
		defaultApikeysModel: newApikeysModel(conn, c, opts...),
	}
}

// CountByUserId 根据用户ID查询API密钥总数
func (m *customApikeysModel) CountByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64
	query := "SELECT COUNT(*) FROM apikeys WHERE user_id = ?"

	err := m.QueryRowNoCacheCtx(ctx, &count, query, userId)
	return count, err
}

// ListByUserId 根据用户ID分页查询API密钥列表
func (m *customApikeysModel) ListByUserId(ctx context.Context, userId int64, limit, offset int32) ([]*Apikeys, error) {
	var apiKeys []*Apikeys
	query := "SELECT id, user_id, name, api_key, created_at, status FROM apikeys WHERE user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?"

	err := m.QueryRowsNoCacheCtx(ctx, &apiKeys, query, userId, limit, offset)
	return apiKeys, err
}

// FindOneByIdAndUserId 根据ID和用户ID查询单个API密钥
func (m *customApikeysModel) FindOneByIdAndUserId(ctx context.Context, id string, userId int64) (*Apikeys, error) {
	var apikey Apikeys
	query := "SELECT id, user_id, name, api_key, created_at, status FROM apikeys WHERE id = ? AND user_id = ?"

	err := m.QueryRowNoCacheCtx(ctx, &apikey, query, id, userId)
	if err != nil {
		return nil, err
	}
	return &apikey, nil
}
