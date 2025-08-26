package model

import (
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ChatInput struct {
	Ai         string `json:"ai" v:"required#AI类型不能为空" dc:"AI类型"`
	Model      string `json:"model" v:"required#模型不能为空" dc:"模型"`
	Prompt     string `json:"prompt" dc:"临时提示词"`
	Message    string `json:"message" v:"required#消息不能为空" dc:"消息"`
	DatabaseId int    `json:"databaseId" v:"required#数据库配置ID不能为空" dc:"数据库配置ID"`
}

type ChatOutDataItem struct {
	Event      string `json:"event"`
	Role       string `json:"role,omitempty"`
	Content    string `json:"content,omitempty"`
	Data       string `json:"data,omitempty"`
	CreateTime int64  `json:"createTime"`
}

func GenChatOutDataItem(event, role string, data any) (ChatOutDataItem, error) {
	contextData := ""
	if !g.IsEmpty(data) {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return ChatOutDataItem{}, err
		}
		contextData = string(jsonData)
	}
	return ChatOutDataItem{
		Event:      event,
		Role:       role,
		Data:       contextData,
		CreateTime: gtime.TimestampMilli(),
	}, nil
}

func SendChatOutDataItem(event, role string, data any, respChan chan any) (err error) {
	item, err := GenChatOutDataItem(event, role, data)
	if err != nil {
		return
	}
	respChan <- item
	return
}
