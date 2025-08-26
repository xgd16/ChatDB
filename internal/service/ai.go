// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"ai-chat-sql/internal/model"
	"context"

	einoModel "github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

type (
	IAI interface {
		// GetChatModel 获取聊天模型
		GetChatModel(ai string, model string) (chatModel einoModel.ToolCallingChatModel, err error)
		// GetChatModeListJson 获取聊天模型列表
		GetChatModeListJson(ctx context.Context, ai string) (json string, err error)
	}
	IAiChat interface {
		AiChatStreamOut(ctx context.Context, respChan chan any, stream *schema.StreamReader[*schema.Message], cancel context.CancelFunc)
		AiChatHeartbeat(ctx context.Context, respChan chan any)
		// Chat 聊天
		Chat(ctx context.Context, in model.ChatInput) (respChan chan any, err error)
	}
)

var (
	localAI     IAI
	localAiChat IAiChat
)

func AI() IAI {
	if localAI == nil {
		panic("implement not found for interface IAI, forgot register?")
	}
	return localAI
}

func RegisterAI(i IAI) {
	localAI = i
}

func AiChat() IAiChat {
	if localAiChat == nil {
		panic("implement not found for interface IAiChat, forgot register?")
	}
	return localAiChat
}

func RegisterAiChat(i IAiChat) {
	localAiChat = i
}
