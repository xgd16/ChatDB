// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// DatabaseConf is the golang structure for table database_conf.
type DatabaseConf struct {
	DatabaseId int    `json:"databaseId" orm:"database_id" description:""` //
	DbName     string `json:"dbName"     orm:"db_name"     description:""` //
	UserName   string `json:"userName"   orm:"user_name"   description:""` //
	Password   string `json:"password"   orm:"password"    description:""` //
	Host       string `json:"host"       orm:"host"        description:""` //
	Port       int    `json:"port"       orm:"port"        description:""` //
	CreateTime int    `json:"createTime" orm:"create_time" description:""` //
	UpdateTime int    `json:"updateTime" orm:"update_time" description:""` //
	DbType     string `json:"dbType"     orm:"db_type"     description:""` //
}
