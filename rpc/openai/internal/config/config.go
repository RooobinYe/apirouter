package config

import (
	"os"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"gopkg.in/yaml.v2"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource   string
	Cache        cache.CacheConf
	Apikey       zrpc.RpcClientConf
	OpenAIAPIKey string // 缓存的API密钥，启动时读取
}

// SecretsConfig 密钥配置结构
type SecretsConfig struct {
	OpenAIAPIKey string `yaml:"OpenAIAPIKey"`
}

// LoadSecrets 启动时加载密钥配置
func (c *Config) LoadSecrets() error {
	data, err := os.ReadFile("rpc/openai/etc/secrets.yaml")
	if err != nil {
		return err
	}

	var secrets SecretsConfig
	if err := yaml.Unmarshal(data, &secrets); err != nil {
		return err
	}

	c.OpenAIAPIKey = secrets.OpenAIAPIKey
	return nil
}

// GetOpenAIAPIKey 获取缓存的OpenAI API Key
func (c *Config) GetOpenAIAPIKey() string {
	return c.OpenAIAPIKey
}
