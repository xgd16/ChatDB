package model

// DatabaseConfigItem 用于数据库配置列表响应，不包含密码等敏感信息
type DatabaseConfigItem struct {
	DatabaseId int    `json:"databaseId" dc:"数据库配置ID"`
	DbName     string `json:"dbName" dc:"数据库名称"`
	UserName   string `json:"userName" dc:"用户名"`
	Host       string `json:"host" dc:"主机地址"`
	Port       int    `json:"port" dc:"端口号"`
	DbType     string `json:"dbType" dc:"数据库类型，如: mysql, postgres, sqlite"`
	CreateTime int    `json:"createTime" dc:"创建时间"`
	UpdateTime int    `json:"updateTime" dc:"更新时间"`
}
