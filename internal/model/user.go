package model

type UserLoginInput struct {
	Username string `json:"username" v:"required#请填写账号" sm:"用户账号" dc:"用户账号"`
	Password string `json:"password" v:"required#请填写密码" sm:"用户密码" dc:"用户密码"`
}

type UserRegisterInput struct {
	Key      string `json:"key" v:"required#请填写注册key" sm:"注册key" dc:"注册key"`
	Username string `json:"username" v:"required#请填写账号" sm:"用户账号" dc:"用户账号"`
	Password string `json:"password" v:"required#请填写密码|min:6#密码不能小于6位" sm:"用户密码" dc:"用户密码"`
}

type User struct {
	UserId       int    `json:"userId"       orm:"user_id"        description:"用户ID"`   //
	Username     string `json:"username"     orm:"username"       description:"用户名"`    //
	RuleLevel    int    `json:"ruleLevel"    orm:"rule_level"     description:"用户权限等级"` //
	LastLoginTme int    `json:"lastLoginTme" orm:"last_login_tme" description:"最后登录时间"` //
	CreateTime   int    `json:"createTime"   orm:"create_time"    description:"创建时间"`   //
}
