package jwt

import (
	"ai-chat-sql/internal/model"
	"ai-chat-sql/internal/service"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v5"
)

type sJwt struct {
}

func init() {
	service.RegisterJwt(NewJwt())
}

func NewJwt() *sJwt {
	return &sJwt{}
}

// GenToken 生成token
func (s *sJwt) GenToken(ctx context.Context, in *model.JWTGenTokenInput) (out *model.JWTGenTokenOutput, err error) {
	out = &model.JWTGenTokenOutput{}
	option, err := service.Config().GetJwtOptions(ctx, in.Subject)
	if err != nil {
		return
	}
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(option.Expire) * time.Hour)), // 过期时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                                               // 签发时间
		Issuer:    option.Issuer,                                                                // 签发者
		Subject:   option.Subject,                                                               // 主题
		ID:        grand.S(12, false),                                                           // 令牌ID
	}

	customClaims := struct {
		jwt.RegisteredClaims
		Id int64 `json:"id"`
	}{
		RegisteredClaims: claims,
		Id:               in.Id,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	token, err := t.SignedString([]byte(option.Secret))
	if err != nil {
		return
	}
	out.Token = token
	out.Expire = claims.ExpiresAt.UnixMilli()
	return
}

func (s *sJwt) VerifyToken(ctx context.Context, in *model.JWTVerifyTokenInput) (out *model.JWTVerifyTokenOutput, valid bool, err error) {
	out = &model.JWTVerifyTokenOutput{}
	option, err := service.Config().GetJwtOptions(ctx, in.Subject)
	if err != nil {
		return
	}
	// 解析并验证令牌
	parsedToken, err := jwt.Parse(in.Token, func(t *jwt.Token) (interface{}, error) {
		// 检查签名算法
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(option.Secret), nil
	})
	if err != nil {
		return
	}

	// 检查令牌是否有效
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		out = &model.JWTVerifyTokenOutput{
			Expire: int(claims["exp"].(float64)),
			Id:     int(claims["id"].(float64)),
			JTI:    claims["jti"].(string),
		}
		valid = true
	}
	return
}
