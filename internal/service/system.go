// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ISystemInit interface {
		// Init 初始化
		Init(ctx context.Context) (err error)
		// CreatePrivateKey 创建私钥
		CreatePrivateKey(ctx context.Context) (err error)
	}
	IPrivateKey interface {
		// Verify 验证私钥
		Verify(ctx context.Context, privateKey string) (err error)
	}
)

var (
	localSystemInit ISystemInit
	localPrivateKey IPrivateKey
)

func SystemInit() ISystemInit {
	if localSystemInit == nil {
		panic("implement not found for interface ISystemInit, forgot register?")
	}
	return localSystemInit
}

func RegisterSystemInit(i ISystemInit) {
	localSystemInit = i
}

func PrivateKey() IPrivateKey {
	if localPrivateKey == nil {
		panic("implement not found for interface IPrivateKey, forgot register?")
	}
	return localPrivateKey
}

func RegisterPrivateKey(i IPrivateKey) {
	localPrivateKey = i
}
