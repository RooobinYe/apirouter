// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3
// Source: user.proto

package userclient

import (
	"context"

	"apirouter/rpc/user/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserInfoRequest    = user.GetUserInfoRequest
	GetUserInfoResponse   = user.GetUserInfoResponse
	LoginData             = user.LoginData
	LoginRequest          = user.LoginRequest
	LoginResponse         = user.LoginResponse
	RegisterData          = user.RegisterData
	RegisterRequest       = user.RegisterRequest
	RegisterResponse      = user.RegisterResponse
	TokenData             = user.TokenData
	UserInfo              = user.UserInfo
	ValidateTokenRequest  = user.ValidateTokenRequest
	ValidateTokenResponse = user.ValidateTokenResponse

	User interface {
		// 用户注册
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		// 用户登录
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		// 验证Token（供AuthMiddleware使用）
		ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
		// 获取用户信息（内部服务使用）
		GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// 用户注册
func (m *defaultUser) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

// 用户登录
func (m *defaultUser) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 验证Token（供AuthMiddleware使用）
func (m *defaultUser) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.ValidateToken(ctx, in, opts...)
}

// 获取用户信息（内部服务使用）
func (m *defaultUser) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}
