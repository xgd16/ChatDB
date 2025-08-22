// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DatabaseConf is the golang structure of table database_conf for DAO operations like Where/Data.
type DatabaseConf struct {
	g.Meta     `orm:"table:database_conf, do:true"`
	DatabaseId interface{} //
	DbName     interface{} //
	UserName   interface{} //
	Password   interface{} //
	Host       interface{} //
	Port       interface{} //
	CreateTime interface{} //
	UpdateTime interface{} //
	DbType     interface{} //
}
