// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/cloudwego/eino/components/model"
)

type (
	IAI interface {
		// GetChatModel 获取聊天模型
		GetChatModel(ai string) (chatModel model.BaseChatModel, err error)
		// GetChatModeListJson 获取聊天模型列表
		GetChatModeListJson(ctx context.Context, ai string) (json string, err error)
	}
)

var (
	localAI IAI
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
