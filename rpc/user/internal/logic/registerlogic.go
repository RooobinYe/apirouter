package logic

import (
	"context"
	"time"

	"apirouter/rpc/model"
	"apirouter/rpc/user/internal/svc"
	"apirouter/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 参数验证
	if in.Username == "" || in.Password == "" || in.Email == "" {
		return &user.RegisterResponse{
			Code:    400,
			Message: "用户名、密码和邮箱不能为空",
		}, nil
	}

	// 检查用户名是否已存在
	_, err := l.svcCtx.UsersModel.FindOneByUsername(l.ctx, in.Username)
	if err == nil {
		return &user.RegisterResponse{
			Code:    400,
			Message: "用户名已存在",
		}, nil
	}
	if err != model.ErrNotFound {
		l.Errorf("Failed to check username: %v", err)
		return &user.RegisterResponse{
			Code:    500,
			Message: "检查用户名失败",
		}, nil
	}

	// 检查邮箱是否已存在
	_, err = l.svcCtx.UsersModel.FindOneByEmail(l.ctx, in.Email)
	if err == nil {
		return &user.RegisterResponse{
			Code:    400,
			Message: "邮箱已存在",
		}, nil
	}
	if err != model.ErrNotFound {
		l.Errorf("Failed to check email: %v", err)
		return &user.RegisterResponse{
			Code:    500,
			Message: "检查邮箱失败",
		}, nil
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Errorf("Failed to hash password: %v", err)
		return &user.RegisterResponse{
			Code:    500,
			Message: "密码加密失败",
		}, nil
	}

	// 创建用户数据
	userData := &model.Users{
		Username:  in.Username,
		Password:  string(hashedPassword),
		Email:     in.Email,
		Status:    1, // active
		CreatedAt: time.Now().Unix(),
	}

	// 插入数据库
	result, err := l.svcCtx.UsersModel.Insert(l.ctx, userData)
	if err != nil {
		l.Errorf("Failed to insert user: %v", err)
		return &user.RegisterResponse{
			Code:    500,
			Message: "注册用户失败",
		}, nil
	}

	// 获取新用户的ID
	userId, err := result.LastInsertId()
	if err != nil {
		l.Errorf("Failed to get user id: %v", err)
		return &user.RegisterResponse{
			Code:    500,
			Message: "获取用户ID失败",
		}, nil
	}

	// 构造响应数据
	registerData := &user.RegisterData{
		UserId:   userId,
		Username: in.Username,
		Email:    in.Email,
	}

	return &user.RegisterResponse{
		Code:    200,
		Message: "注册成功",
		Data:    registerData,
	}, nil
}
