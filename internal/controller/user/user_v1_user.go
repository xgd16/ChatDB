package user

import (
	"context"

	v1 "ai-chat-sql/api/user/v1"
	"ai-chat-sql/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	user, out, err := service.User().Login(ctx, req.UserLoginInput)
	if err != nil {
		return
	}
	res = &v1.UserLoginRes{}
	res.JWTGenTokenOutput = *out
	err = gconv.Scan(user, &res.User)
	return
}
func (c *ControllerV1) UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	res = &v1.UserRegisterRes{}

	userId, err := service.User().Register(ctx, true, req.UserRegisterInput)
	if err != nil {
		return
	}
	out, err := service.User().GenJwtTokenByUserId(ctx, userId)
	if err != nil {
		return
	}

	res.JWTGenTokenOutput = *out
	return
}
