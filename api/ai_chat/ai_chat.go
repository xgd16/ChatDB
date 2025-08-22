// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ai_chat

import (
	"context"

	"ai-chat-sql/api/ai_chat/v1"
)

type IAiChatV1 interface {
	Chat(ctx context.Context, req *v1.ChatReq) (res *v1.ChatRes, err error)
	SetDataBaseConfig(ctx context.Context, req *v1.SetDataBaseConfigReq) (res *v1.SetDataBaseConfigRes, err error)
	UpdateDataBaseConfig(ctx context.Context, req *v1.UpdateDataBaseConfigReq) (res *v1.UpdateDataBaseConfigRes, err error)
	DeleteDataBaseConfig(ctx context.Context, req *v1.DeleteDataBaseConfigReq) (res *v1.DeleteDataBaseConfigRes, err error)
	GetDataBaseConfigList(ctx context.Context, req *v1.GetDataBaseConfigListReq) (res *v1.GetDataBaseConfigListRes, err error)
	GetAiModelList(ctx context.Context, req *v1.GetAiModelListReq) (res *v1.GetAiModelListRes, err error)
}
