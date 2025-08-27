package user

import (
	"context"

	v1 "ai-chat-sql/api/user/v1"
	"ai-chat-sql/internal/service"
)

func (c *ControllerV1) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {

	return
}
func (c *ControllerV1) UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	_, err = service.User().Register(ctx, true, req.UserRegisterInput)
	if err != nil {
		return
	}
	return
}
