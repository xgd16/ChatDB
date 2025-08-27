// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"ai-chat-sql/internal/model"
	"ai-chat-sql/internal/model/entity"
	"context"
)

type (
	IUser interface {
		// Login 登录
		Login(ctx context.Context, in model.UserLoginInput) (user *entity.User, out *model.JWTGenTokenOutput, err error)
		// Register 注册用户
		Register(ctx context.Context, verify bool, in model.UserRegisterInput) (userId int64, err error)
		// GenJwtTokenByUserId 根据用户ID生成JWT
		GenJwtTokenByUserId(ctx context.Context, userId int64) (out *model.JWTGenTokenOutput, err error)
		// GenPassword 生成密码
		GenPassword(password string, code int) (out string, err error)
		// GetUserVerifyCode 获取用户验证码
		GetUserVerifyCode() (code int)
		// UsernameIsExist 用户名是否存在
		UsernameIsExist(ctx context.Context, username string) (exist bool, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
