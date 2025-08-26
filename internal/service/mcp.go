// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"ai-chat-sql/internal/model"
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type (
	IMcpHandler interface {
		GetList() []model.McpReg
		GetMcpFn(item *model.McpReg) server.ToolHandlerFunc
	}
	IMcpTool interface {
		// ExecSql 执行SQL
		ExecSql(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		GetNowTime(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// TimestampToDateTime 时间戳转换为日期时间
		TimestampToDateTime(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
	}
)

var (
	localMcpHandler IMcpHandler
	localMcpTool    IMcpTool
)

func McpHandler() IMcpHandler {
	if localMcpHandler == nil {
		panic("implement not found for interface IMcpHandler, forgot register?")
	}
	return localMcpHandler
}

func RegisterMcpHandler(i IMcpHandler) {
	localMcpHandler = i
}

func McpTool() IMcpTool {
	if localMcpTool == nil {
		panic("implement not found for interface IMcpTool, forgot register?")
	}
	return localMcpTool
}

func RegisterMcpTool(i IMcpTool) {
	localMcpTool = i
}
