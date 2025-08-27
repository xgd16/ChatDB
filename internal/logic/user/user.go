package user

import (
	"ai-chat-sql/internal/code"
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/dao"
	"ai-chat-sql/internal/model"
	"ai-chat-sql/internal/model/entity"
	"ai-chat-sql/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type sUser struct {
}

func NewUser() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(NewUser())
}

// Login 登录
func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) (user *entity.User, out *model.JWTGenTokenOutput, err error) {
	// 验证用户名是否存在
	user, err = dao.User.GetUserInfoByUsername(ctx, in.Username)
	if err != nil {
		return
	}
	// 验证密码是否正确
	password, err := s.GenPassword(in.Password, user.Verify)
	if err != nil {
		return
	}
	if user.Password != password {
		err = code.ToError(code.LoginFailMessage)
		return
	}
	out, err = s.GenJwtTokenByUserId(ctx, int64(user.UserId))
	return
}

// Register 注册用户
func (s *sUser) Register(ctx context.Context, verify bool, in model.UserRegisterInput) (userId int64, err error) {
	// 验证私钥
	if verify {
		if err = service.PrivateKey().Verify(ctx, in.Key); err != nil {
			return
		}
	}
	// 验证用户名是否存在
	exist, err := s.UsernameIsExist(ctx, in.Username)
	if err != nil {
		return
	}
	if exist {
		err = code.ToError(code.UsernameExist)
		return
	}
	// 生成验证码
	code := s.GetUserVerifyCode()
	// 生成密码
	password, err := s.GenPassword(in.Password, code)
	if err != nil {
		return
	}
	// 注册用户
	userId, err = dao.User.Ctx(ctx).Data(g.Map{
		"username":   in.Username,
		"password":   password,
		"verify":     code,
		"rule_level": 1,
	}).InsertAndGetId()
	return
}

// GenJwtTokenByUserId 根据用户ID生成JWT
func (s *sUser) GenJwtTokenByUserId(ctx context.Context, userId int64) (out *model.JWTGenTokenOutput, err error) {
	out, err = service.Jwt().GenToken(ctx, &model.JWTGenTokenInput{
		Id:      userId,
		Subject: consts.JwtSubjectUser,
	})
	if err != nil {
		return
	}
	return
}

// GenPassword 生成密码
func (s *sUser) GenPassword(password string, code int) (out string, err error) {
	out, err = gmd5.Encrypt(fmt.Sprintf("ai-chat:%s:%d", password, code))
	if err != nil {
		return
	}
	return
}

// GetUserVerifyCode 获取用户验证码
func (s *sUser) GetUserVerifyCode() (code int) {
	code = grand.N(10000, 99999)
	return
}

// UsernameIsExist 用户名是否存在
func (s *sUser) UsernameIsExist(ctx context.Context, username string) (exist bool, err error) {
	count, err := dao.User.Ctx(ctx).Where("username = ?", username).Count()
	if err != nil {
		return
	}
	return count > 0, nil
}
