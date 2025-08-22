package config

import (
	"ai-chat-sql/internal/dao"
	"ai-chat-sql/internal/model/entity"
	"ai-chat-sql/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
)

type sConfig struct {
}

func init() {
	service.RegisterConfig(NewConfig())
}

func NewConfig() *sConfig {
	return &sConfig{}
}

// GetDataBase 获取数据库连接
func (s *sConfig) GetDataBase(ctx context.Context, databaseId int) (db gdb.DB, err error) {
	link, err := s.GenDataBaseLink(ctx, databaseId)
	if err != nil {
		return
	}
	db, err = gdb.New(gdb.ConfigNode{
		Link: link,
	})
	if err != nil {
		return
	}
	return
}

// GenDataBaseLink 生成数据库连接
//
// 格式
// mysql:root:12345678@tcp(127.0.0.1:3306)/test?loc=Local&parseTime=true
// pgsql:root:12345678@tcp(127.0.0.1:5432)/test
// sqlite::@file(/var/data/db.sqlite3)
func (s *sConfig) GenDataBaseLink(ctx context.Context, databaseId int) (link string, err error) {
	// 根据数据库ID查询配置信息
	var config *entity.DatabaseConf
	err = dao.DatabaseConf.Ctx(ctx).Where("database_id = ?", databaseId).Scan(&config)
	if err != nil {
		return "", fmt.Errorf("查询数据库配置失败: %w", err)
	}
	if config == nil {
		return "", fmt.Errorf("数据库配置不存在")
	}

	// 根据数据库类型生成连接字符串
	switch config.DbType {
	case "mysql":
		link = fmt.Sprintf("mysql:%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			config.UserName,
			config.Password,
			config.Host,
			config.Port,
			config.DbName,
		)
	case "pgsql":
		link = fmt.Sprintf("pgsql:%s:%s@tcp(%s:%d)/%s",
			config.UserName,
			config.Password,
			config.Host,
			config.Port,
			config.DbName,
		)
	case "sqlite":
		link = fmt.Sprintf("sqlite::@file(%s)", config.DbName)
	default:
		return "", fmt.Errorf("不支持的数据库类型: %s", config.DbType)
	}

	return link, nil
}
