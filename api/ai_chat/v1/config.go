package v1

import (
	"ai-chat-sql/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type SetDataBaseConfigReq struct {
	g.Meta   `path:"/config/db/setConfig" method:"post" tags:"V1/配置/数据库" sm:"设置数据库配置" dc:"设置数据库配置"`
	DbName   string `json:"dbName" v:"required#数据库名称不能为空" dc:"数据库名称"`
	UserName string `json:"userName" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Host     string `json:"host" v:"required#主机地址不能为空" dc:"主机地址"`
	Port     int    `json:"port" v:"required#端口号不能为空" dc:"端口号"`
	DbType   string `json:"dbType" v:"required#数据库类型不能为空" dc:"数据库类型，如: mysql, postgres, sqlite"`
}

type SetDataBaseConfigRes struct {
}

type UpdateDataBaseConfigReq struct {
	g.Meta     `path:"/config/db/updateConfig" method:"put" tags:"V1/配置/数据库" sm:"更新数据库配置" dc:"更新数据库配置"`
	DatabaseId int    `json:"databaseId" v:"required#数据库配置ID不能为空" dc:"数据库配置ID"`
	DbName     string `json:"dbName" v:"required#数据库名称不能为空" dc:"数据库名称"`
	UserName   string `json:"userName" v:"required#用户名不能为空" dc:"用户名"`
	Password   string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Host       string `json:"host" v:"required#主机地址不能为空" dc:"主机地址"`
	Port       int    `json:"port" v:"required#端口号不能为空" dc:"端口号"`
	DbType     string `json:"dbType" v:"required#数据库类型不能为空" dc:"数据库类型，如: mysql, postgres, sqlite"`
}

type UpdateDataBaseConfigRes struct {
}

type DeleteDataBaseConfigReq struct {
	g.Meta     `path:"/config/db/deleteConfig" method:"delete" tags:"V1/配置/数据库" sm:"删除数据库配置" dc:"删除数据库配置"`
	DatabaseId int `json:"databaseId" v:"required#数据库配置ID不能为空" dc:"数据库配置ID"`
}

type DeleteDataBaseConfigRes struct {
}

type GetDataBaseConfigListReq struct {
	g.Meta `path:"/config/db/listConfig" method:"get" tags:"V1/配置/数据库" sm:"获取数据库配置列表" dc:"获取数据库配置列表"`
	Page   int `json:"page" d:"1" dc:"页码"`
	Size   int `json:"size" d:"10" dc:"每页数量"`
}

type GetDataBaseConfigListRes struct {
	List  []*SetDataBaseConfigReq `json:"list" dc:"数据库配置列表"`
	Total int                     `json:"total" dc:"总数"`
}

type GetAiModelListReq struct {
	g.Meta `path:"/config/ai/model/list" method:"get" tags:"V1/配置" sm:"获取AI模型列表" dc:"获取AI模型列表 (只有Open Ai 协议支持)"`
}

type GetAiModelListRes struct {
	List  []*model.AIModelItem `json:"list" dc:"AI模型列表"`
	Total int                  `json:"total" dc:"总数"`
}
