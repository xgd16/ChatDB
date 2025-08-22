package v1

import "github.com/gogf/gf/v2/frame/g"

type ChatReq struct {
	g.Meta     `path:"/chats" method:"post" tags:"V1/聊天" sm:"聊天" dc:"聊天"`
	Prompt     string `json:"prompt" v:"required#提示不能为空" dc:"提示"`
	DatabaseId int    `json:"databaseId" v:"required#数据库配置ID不能为空" dc:"数据库配置ID"`
}

type ChatRes struct {
}
