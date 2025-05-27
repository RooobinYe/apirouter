package logic

import (
	"context"
	"time"

	"apirouter/rpc/model"
	"apirouter/rpc/user/internal/svc"
	"apirouter/rpc/user/user"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 参数验证
	if in.Username == "" || in.Password == "" {
		return &user.LoginResponse{
			Code:    400,
			Message: "用户名和密码不能为空",
		}, nil
	}

	// 查找用户
	userData, err := l.svcCtx.UsersModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			return &user.LoginResponse{
				Code:    401,
				Message: "用户名或密码错误",
			}, nil
		}
		l.Errorf("Failed to find user: %v", err)
		return &user.LoginResponse{
			Code:    500,
			Message: "登录失败",
		}, nil
	}

	// 检查用户状态
	if userData.Status != 1 {
		return &user.LoginResponse{
			Code:    401,
			Message: "账户已被禁用",
		}, nil
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(in.Password))
	if err != nil {
		return &user.LoginResponse{
			Code:    401,
			Message: "用户名或密码错误",
		}, nil
	}

	// 生成JWT Token
	now := time.Now()
	expiresAt := now.Add(time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire) * time.Second)

	claims := jwt.MapClaims{
		"user_id":  userData.Id,
		"username": userData.Username,
		"exp":      expiresAt.Unix(),
		"iat":      now.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
	if err != nil {
		l.Errorf("Failed to generate token: %v", err)
		return &user.LoginResponse{
			Code:    500,
			Message: "生成访问令牌失败",
		}, nil
	}

	// 构造响应数据
	loginData := &user.LoginData{
		UserId:      userData.Id,
		Username:    userData.Username,
		Email:       userData.Email,
		AccessToken: accessToken,
	}

	return &user.LoginResponse{
		Code:    200,
		Message: "登录成功",
		Data:    loginData,
	}, nil
}
