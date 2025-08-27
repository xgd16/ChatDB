package middleware

import (
	"ai-chat-sql/internal/code"
	"ai-chat-sql/internal/model"
	"ai-chat-sql/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

// JwtAuth 校验jwt
func (s *sMiddleware) JwtAuth(subject string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		// 获取当前路由处理函数
		handler := r.GetServeHandler()
		// 检查是否有noAuth元数据，如果有且值为true，则跳过验证
		noAuth := handler.GetMetaTag("noAuth")
		if noAuth == "true" {
			r.Middleware.Next()
			return
		}

		ctx := r.GetCtx()
		authorization := r.GetHeader("authorization")
		token := gstr.TrimLeft(authorization, "Bearer ")
		data, v, err := service.Jwt().VerifyToken(ctx, &model.JWTVerifyTokenInput{
			Token:   token,
			Subject: subject,
		})

		if err != nil {
			r.Response.WriteJsonExit(g.Map{
				"code":        code.FailedToRequest.Code(),
				"message":     code.FailedToRequest.Message(),
				"serviceTime": gtime.Now().UnixMilli(),
			})
		}
		if !v {
			r.Response.WriteJsonExit(g.Map{
				"code":        code.InvalidToken.Code(),
				"message":     code.InvalidToken.Message(),
				"serviceTime": gtime.Now().UnixMilli(),
			})
		}
		r.SetCtx(context.WithValue(ctx, model.UserGroup{}, data.Id))
		r.Middleware.Next()
	}
}
