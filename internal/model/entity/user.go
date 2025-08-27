// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	UserId       int    `json:"userId"       orm:"user_id"        description:""` //
	Username     string `json:"username"     orm:"username"       description:""` //
	Password     string `json:"password"     orm:"password"       description:""` //
	Verify       int    `json:"verify"       orm:"verify"         description:""` //
	RuleLevel    int    `json:"ruleLevel"    orm:"rule_level"     description:""` //
	LastLoginTme int    `json:"lastLoginTme" orm:"last_login_tme" description:""` //
	CreateTime   int    `json:"createTime"   orm:"create_time"    description:""` //
	UpdateTime   int    `json:"updateTime"   orm:"update_time"    description:""` //
}
