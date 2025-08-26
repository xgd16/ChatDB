// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"ai-chat-sql/internal/model"
	"context"
)

type (
	IPrompt interface {
		// GetList 获取提示词列表
		GetList(ctx context.Context) (out []*model.PromptGetListOutputItem, err error)
		// CleanCache 清理缓存
		CleanCache()
		// GetPrompt 获取提示词
		GetPrompt(ctx context.Context, fileName string) (out *model.PromptGetListOutputItem, err error)
	}
)

var (
	localPrompt IPrompt
)

func Prompt() IPrompt {
	if localPrompt == nil {
		panic("implement not found for interface IPrompt, forgot register?")
	}
	return localPrompt
}

func RegisterPrompt(i IPrompt) {
	localPrompt = i
}
