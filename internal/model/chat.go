package model

type ChatInput struct {
	Ai         string `json:"ai" v:"required#AI类型不能为空" dc:"AI类型"`
	Model      string `json:"model" v:"required#模型不能为空" dc:"模型"`
	Prompt     string `json:"prompt" dc:"临时提示词"`
	Message    string `json:"message" v:"required#消息不能为空" dc:"消息"`
	DatabaseId int    `json:"databaseId" v:"required#数据库配置ID不能为空" dc:"数据库配置ID"`
}
