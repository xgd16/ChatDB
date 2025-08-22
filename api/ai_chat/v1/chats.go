package v1

import (
	"ai-chat-sql/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ChatReq struct {
	g.Meta `path:"/chats" method:"post" tags:"V1/聊天" sm:"聊天" dc:"聊天"`
	model.ChatInput
}

type ChatRes struct {
	g.Meta `mime:"text/event-stream"`
}
