// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"ai-chat-sql/api/user/v1"
)

type IUserV1 interface {
	UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error)
	UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error)
}
