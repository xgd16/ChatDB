package ai_chat

import (
	"context"

	v1 "ai-chat-sql/api/ai_chat/v1"
	"ai-chat-sql/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Chat(ctx context.Context, req *v1.ChatReq) (res *v1.ChatRes, err error) {
	db, err := service.Config().GetDataBase(ctx, req.DatabaseId)
	if err != nil {
		return
	}
	g.DumpWithType(db.Query(ctx, "SHOW TABLES"))
	return
}
