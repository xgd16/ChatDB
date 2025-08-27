package v1

import (
	"ai-chat-sql/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginReq struct {
	g.Meta `path:"/user/login" method:"post" tags:"V1/用户" sm:"登录" dc:"登录" noAuth:"true"`
	model.UserLoginInput
}

type UserLoginRes struct {
	model.JWTGenTokenOutput
	model.User
}

type UserRegisterReq struct {
	g.Meta `path:"/user/register" method:"post" tags:"V1/用户" sm:"注册" dc:"注册" noAuth:"true"`
	model.UserRegisterInput
}

type UserRegisterRes struct {
	model.JWTGenTokenOutput
}
