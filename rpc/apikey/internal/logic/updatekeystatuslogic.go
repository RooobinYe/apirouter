package logic

import (
	"context"

	"apirouter/rpc/apikey/apikey"
	"apirouter/rpc/apikey/internal/svc"
	"apirouter/rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateKeyStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateKeyStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateKeyStatusLogic {
	return &UpdateKeyStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新API密钥状态（内部服务使用）
func (l *UpdateKeyStatusLogic) UpdateKeyStatus(in *apikey.UpdateKeyStatusRequest) (*apikey.UpdateKeyStatusResponse, error) {
	// 验证状态值
	validStatuses := map[string]bool{
		"active":   true,
		"inactive": true,
		"expired":  true,
	}
	
	if !validStatuses[in.Status] {
		return &apikey.UpdateKeyStatusResponse{
			Code:    400,
			Message: "无效的状态值",
		}, nil
	}

	// 检查API密钥是否存在且属于指定用户
	_, err := l.svcCtx.ApiKeysModel.FindOneByIdAndUserId(l.ctx, in.Id, in.UserId)
	if err != nil {
		l.Errorf("Failed to find api key: %v", err)
		return &apikey.UpdateKeyStatusResponse{
			Code:    404,
			Message: "API密钥不存在",
		}, nil
	}

	// 构造更新的数据
	data := &model.Apikeys{
		Id:     in.Id,
		Status: in.Status,
	}

	// 更新API密钥状态
	err = l.svcCtx.ApiKeysModel.Update(l.ctx, data)
	if err != nil {
		l.Errorf("Failed to update api key status: %v", err)
		return &apikey.UpdateKeyStatusResponse{
			Code:    500,
			Message: "更新API密钥状态失败",
		}, nil
	}

	return &apikey.UpdateKeyStatusResponse{
		Code:    200,
		Message: "API密钥状态更新成功",
		Data: &apikey.SuccessData{
			Message: "状态更新成功",
		},
	}, nil
}
