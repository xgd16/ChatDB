// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"ai-chat-sql/internal/model"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IConfig interface {
		// GetDataBase 获取数据库连接
		GetDataBase(ctx context.Context, databaseId int) (db gdb.DB, err error)
		// GenDataBaseLink 生成数据库连接
		//
		// 格式
		// mysql:root:12345678@tcp(127.0.0.1:3306)/test?loc=Local&parseTime=true
		// pgsql:root:12345678@tcp(127.0.0.1:5432)/test
		// sqlite::@file(/var/data/db.sqlite3)
		GenDataBaseLink(ctx context.Context, databaseId int) (link string, err error)
		// GetJwtOptions 获取JWT配置
		GetJwtOptions(ctx context.Context, secret string) (option *model.JwtOption, err error)
	}
)

var (
	localConfig IConfig
)

func Config() IConfig {
	if localConfig == nil {
		panic("implement not found for interface IConfig, forgot register?")
	}
	return localConfig
}

func RegisterConfig(i IConfig) {
	localConfig = i
}
