package mcp

import (
	"ai-chat-sql/internal/service"
	"context"
	"encoding/json"
	"errors"

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
		return
	}
	jsonStr, err := json.Marshal(sqlOut)
	if err != nil {
		return
	}
	// g.DumpWithType(sqlOut)
	out = mcp.NewToolResultText(string(jsonStr))
	return
}
