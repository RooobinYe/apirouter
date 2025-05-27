package main

import (
	"flag"
	"fmt"
	"os"

	"apirouter/rpc/openai/internal/config"
	"apirouter/rpc/openai/internal/server"
	"apirouter/rpc/openai/internal/svc"
	"apirouter/rpc/openai/openai"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/openai.yaml", "the config file")
var secretsFile = flag.String("s", "etc/secrets.yaml", "the secrets file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 检查 secrets.yaml 文件是否存在
	if _, err := os.Stat(*secretsFile); err == nil {
		// 从 secrets.yaml 加载 OpenAI API Key
		var secrets struct {
			OpenAIAPIKey string
		}
		err := conf.Load(*secretsFile, &secrets)
		if err != nil {
			fmt.Printf("Warning: Failed to load secrets file: %v\n", err)
		} else if secrets.OpenAIAPIKey != "" {
			c.OpenAIAPIKey = secrets.OpenAIAPIKey
			fmt.Println("Successfully loaded OpenAI API Key from secrets file")
		}
	} else {
		fmt.Printf("Warning: Secrets file not found: %s\n", *secretsFile)
	}

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		openai.RegisterOpenAIServer(grpcServer, server.NewOpenAIServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
