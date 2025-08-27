package cmd

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/controller/ai_chat"
	"ai-chat-sql/internal/controller/user"
	"ai-chat-sql/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().HandlerResponse)

				{
					authGroup := group.Clone()
					authGroup.Middleware(service.Middleware().JwtAuth(consts.JwtSubjectUser)).
						Bind(ai_chat.NewV1(), user.NewV1())
				}
			})
			s.Run()
			return nil
		},
	}
)
