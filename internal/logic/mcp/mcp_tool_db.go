package mcp

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/service"
	"ai-chat-sql/utility"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mark3labs/mcp-go/mcp"
)

// ExecSql 执行SQL
func (s *sMcpTool) ExecSql(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error) {
	databaseId := request.GetInt("databaseId", 0)
	if databaseId == 0 {
		err = errors.New("databaseId is required")
		return
	}
	sql := request.GetString("sql", "")
	if sql == "" {
		err = errors.New("sql is required")
		return
	}
	db, err := service.Config().GetDataBase(ctx, databaseId)
	if err != nil {
		return
	}

	// 检查是否启用只读模式
	if consts.Config.DbConfig != nil && consts.Config.DbConfig.Readonly {
		if !isReadOnlySQL(sql) {
			errMsg := "数据库当前处于只读模式，只允许执行查询操作（SELECT语句）"
			consts.Logger.Warning(ctx, errMsg)
			out = mcp.NewToolResultText(errMsg)
			err = nil
			return
		}
	}

	sqlOut, err := db.Query(ctx, sql)
	if err != nil {
		outStr := fmt.Sprintf("数据库执行失败：%s", err.Error())
		consts.Logger.Error(ctx, outStr)
		out = mcp.NewToolResultText(outStr)
		err = nil
		return
	}

	respStr, err := utility.ConvertAnyToMarkdownTable(sqlOut.List())
	if err != nil {
		return
	}

	// 在返回结果中包含执行的 SQL 语句信息
	// 格式：SQL 语句作为前缀，然后是执行结果
	fullResult := fmt.Sprintf("**执行的 SQL：**\n\n```sql\n%s\n```\n\n**执行结果：**\n\n%s", sql, respStr)
	out = mcp.NewToolResultText(fullResult)
	return
}

// GetDatabaseInfo 获取数据库信息
func (s *sMcpTool) GetDatabaseInfo(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error) {
	// 兜底防止 g.DB 在未配置 default 数据源时直接 panic，避免整个服务挂掉
	defer func() {
		if r := recover(); r != nil {
			consts.Logger.Errorf(ctx, "GetDatabaseInfo panic: %+v", r)
			// 友好的错误提示返回给前端 / AI
			out = mcp.NewToolResultText("数据库配置错误：未找到默认数据库配置，请检查服务端 config 中的数据库设置")
			err = nil
		}
	}()

	// 可选的数据库名称参数，未提供时使用默认连接
	dbname := request.GetString("dbname", "")

	// 获取数据库配置信息，先检查 DB 对象是否为 nil
	db := g.DB(dbname)
	if db == nil {
		err = errors.New("数据库连接不存在，请检查数据库配置")
		return
	}

	dbConfig := db.GetConfig()
	if dbConfig == nil {
		err = errors.New("无法获取数据库配置")
		return
	}

	// 构建数据库信息
	dbInfo := g.Map{
		"databaseType": dbConfig.Type,
		"host":         extractHostFromLink(dbConfig.Link),
		"port":         extractPortFromLink(dbConfig.Link),
		"databaseName": extractDatabaseNameFromLink(dbConfig.Link),
		"username":     extractUsernameFromLink(dbConfig.Link),
		"prefix":       dbConfig.Prefix,
		"createdAt":    dbConfig.CreatedAt,
		"updatedAt":    dbConfig.UpdatedAt,
		"debug":        dbConfig.Debug,
	}

	// 测试连接并获取数据库版本信息
	versionQuery := getVersionQuery(dbConfig.Type)
	if versionQuery != "" {
		queryDb := g.DB(dbname)
		if queryDb != nil {
			sqlOut, queryErr := queryDb.Query(ctx, versionQuery)
			if queryErr == nil && sqlOut != nil && len(sqlOut.List()) > 0 {
				versionInfo := sqlOut.List()[0]
				dbInfo["version"] = versionInfo
			}
		}
	}

	// 获取数据库大小（如果支持）
	sizeQuery := getSizeQuery(dbConfig.Type, dbname)
	if sizeQuery != "" {
		queryDb := g.DB(dbname)
		if queryDb != nil {
			sqlOut, queryErr := queryDb.Query(ctx, sizeQuery)
			if queryErr == nil && sqlOut != nil && len(sqlOut.List()) > 0 {
				sizeInfo := sqlOut.List()[0]
				dbInfo["databaseSize"] = sizeInfo
			}
		}
	}

	out = mcp.NewToolResultText(gjson.MustEncodeString(dbInfo))
	return
}

// 从连接字符串中提取主机地址
func extractHostFromLink(link string) string {
	// 简单的解析，实际项目中可能需要更复杂的解析
	// 格式通常是: user:pass@tcp(host:port)/dbname?params
	if link == "" {
		return ""
	}

	// 查找 tcp( 和 ) 之间的内容
	start := strings.Index(link, "tcp(")
	if start == -1 {
		return ""
	}
	start += 4 // 跳过 "tcp("

	end := strings.Index(link[start:], ")")
	if end == -1 {
		return ""
	}

	hostPort := link[start : start+end]
	// 分离主机和端口
	parts := strings.Split(hostPort, ":")
	if len(parts) > 0 {
		return parts[0]
	}
	return hostPort
}

// 从连接字符串中提取端口
func extractPortFromLink(link string) string {
	if link == "" {
		return ""
	}

	start := strings.Index(link, "tcp(")
	if start == -1 {
		return ""
	}
	start += 4

	end := strings.Index(link[start:], ")")
	if end == -1 {
		return ""
	}

	hostPort := link[start : start+end]
	parts := strings.Split(hostPort, ":")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

// 从连接字符串中提取数据库名称
func extractDatabaseNameFromLink(link string) string {
	if link == "" {
		return ""
	}

	// 查找最后一个 / 之后的内容
	lastSlash := strings.LastIndex(link, "/")
	if lastSlash == -1 {
		return ""
	}

	dbPart := link[lastSlash+1:]
	// 移除查询参数
	questionMark := strings.Index(dbPart, "?")
	if questionMark != -1 {
		dbPart = dbPart[:questionMark]
	}

	return dbPart
}

// 从连接字符串中提取用户名
func extractUsernameFromLink(link string) string {
	if link == "" {
		return ""
	}

	// 查找 @ 之前的内容
	atIndex := strings.Index(link, "@")
	if atIndex == -1 {
		return ""
	}

	userPass := link[:atIndex]
	// 分离用户名和密码
	colonIndex := strings.Index(userPass, ":")
	if colonIndex != -1 {
		return userPass[:colonIndex]
	}
	return userPass
}

// 根据数据库类型获取版本查询语句
func getVersionQuery(dbType string) string {
	switch dbType {
	case "mysql":
		return "SELECT VERSION() as version"
	case "postgresql", "postgres":
		return "SELECT version() as version"
	case "sqlite":
		return "SELECT sqlite_version() as version"
	case "mssql", "sqlserver":
		return "SELECT @@VERSION as version"
	default:
		return ""
	}
}

// 根据数据库类型获取大小查询语句
func getSizeQuery(dbType, dbname string) string {
	switch dbType {
	case "mysql":
		return fmt.Sprintf("SELECT ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) AS 'Size (MB)' FROM information_schema.tables WHERE table_schema = '%s'", dbname)
	case "postgresql", "postgres":
		return fmt.Sprintf("SELECT pg_size_pretty(pg_database_size('%s')) as size", dbname)
	case "sqlite":
		return "SELECT page_count * page_size as size FROM pragma_page_count(), pragma_page_size()"
	default:
		return ""
	}
}

// isReadOnlySQL 检查SQL语句是否为只读操作
func isReadOnlySQL(sql string) bool {
	// 去除前后空格并转换为大写
	sql = strings.TrimSpace(strings.ToUpper(sql))

	// 检查是否以SELECT开头（包括WITH语句，因为WITH通常用于查询）
	if strings.HasPrefix(sql, "SELECT") || strings.HasPrefix(sql, "WITH") {
		return true
	}

	// 检查是否为SHOW语句（MySQL的SHOW语句是只读的）
	if strings.HasPrefix(sql, "SHOW") {
		return true
	}

	// 检查是否为DESCRIBE或DESC语句
	if strings.HasPrefix(sql, "DESCRIBE") || strings.HasPrefix(sql, "DESC") {
		return true
	}

	// 检查是否为EXPLAIN语句
	if strings.HasPrefix(sql, "EXPLAIN") {
		return true
	}

	// 检查是否为PRAGMA语句（SQLite的PRAGMA语句通常是只读的）
	if strings.HasPrefix(sql, "PRAGMA") {
		return true
	}

	// 其他情况都认为是非只读操作
	return false
}
