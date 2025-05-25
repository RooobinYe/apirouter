package logic

import (
	"apirouter/rpc/user/userclient"
	"context"

	"apirouter/api/internal/svc"
	"apirouter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	rpcResp, err := l.svcCtx.UserClient.Register(l.ctx, &userclient.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		l.Logger.Errorf("register failed: %v", err)
		return &types.RegisterResp{
			Code:    500,
			Message: "服务器内部错误",
		}, nil
	}

	if rpcResp.Code != 200 {
		return &types.RegisterResp{
			Code:    int(rpcResp.Code),
			Message: rpcResp.Message,
		}, nil
	}

	return &types.RegisterResp{
		Code:    200,
		Message: "注册成功",
		Data: types.RegisterData{
			UserId:   rpcResp.Data.UserId,
			Username: rpcResp.Data.Username,
			Email:    rpcResp.Data.Email,
		},
	}, nil
}
