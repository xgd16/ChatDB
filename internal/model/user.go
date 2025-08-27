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
