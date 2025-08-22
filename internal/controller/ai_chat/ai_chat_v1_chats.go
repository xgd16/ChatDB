package ai_chat

import (
	"context"
	"encoding/json"

	v1 "ai-chat-sql/api/ai_chat/v1"
	"ai-chat-sql/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Chat(ctx context.Context, req *v1.ChatReq) (res *v1.ChatRes, err error) {
	respChan, err := service.AiChat().Chat(ctx, req.ChatInput)
	if err != nil {
		return
	}

	r := g.RequestFromCtx(ctx)
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")

	var jsonData []byte
	for v := range respChan {
		switch v.(type) {
		case string:
			r.Response.Writef("%s\n\n", v)
		default:
			jsonData, err = json.Marshal(v)
			if err != nil {
				return
			}
			r.Response.Writef("data: %s\n\n", jsonData)
		}
		r.Response.Flush()
	}
	return
}
