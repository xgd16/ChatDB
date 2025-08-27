// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta       `orm:"table:user, do:true"`
	UserId       interface{} //
	Username     interface{} //
	Password     interface{} //
	Verify       interface{} //
	RuleLevel    interface{} //
	LastLoginTme interface{} //
	CreateTime   interface{} //
	UpdateTime   interface{} //
}
