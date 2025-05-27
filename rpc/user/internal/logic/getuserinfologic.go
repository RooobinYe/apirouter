package logic

import (
	"context"

	"apirouter/rpc/model"
	"apirouter/rpc/user/internal/svc"
	"apirouter/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息（内部服务使用）
func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	// 参数验证
	if in.UserId <= 0 {
		return &user.GetUserInfoResponse{
			Code:    400,
			Message: "用户ID无效",
		}, nil
	}

	// 查询用户信息
	userData, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			return &user.GetUserInfoResponse{
				Code:    404,
				Message: "用户不存在",
			}, nil
		}
		l.Errorf("Failed to find user: %v", err)
		return &user.GetUserInfoResponse{
			Code:    500,
			Message: "获取用户信息失败",
		}, nil
	}

	// 构造用户信息
	userInfo := &user.UserInfo{
		UserId:    userData.Id,
		Username:  userData.Username,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
		Status:    int32(userData.Status),
	}

	return &user.GetUserInfoResponse{
		Code:    200,
		Message: "获取成功",
		Data:    userInfo,
	}, nil
}
