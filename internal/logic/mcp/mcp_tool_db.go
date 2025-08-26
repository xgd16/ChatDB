package mcp

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/service"
	"ai-chat-sql/utility"
	"context"
	"errors"
	"fmt"

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
	// g.DumpWithType(sqlOut)
	out = mcp.NewToolResultText(respStr)
	return
}
