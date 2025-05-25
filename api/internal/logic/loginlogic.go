package logic

import (
	"context"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"
	"apirouter/rpc/user/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.UserClient.Login(l.ctx, &userclient.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return &types.LoginResp{
			Code:    500,
			Message: "登录失败",
		}, err
	}

	if loginResp.Code != 200 {
		return &types.LoginResp{
			Code:    int(loginResp.Code),
			Message: loginResp.Message,
		}, nil
	}

	return &types.LoginResp{
		Code:    200,
		Message: "登录成功",
		Data: types.LoginData{
			UserId:      loginResp.Data.UserId,
			Username:    loginResp.Data.Username,
			Email:       loginResp.Data.Email,
			AccessToken: loginResp.Data.AccessToken,
		},
	}, nil
}
