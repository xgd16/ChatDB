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
		SayHello(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
	}
)

var (
	localMcpHandler IMcpHandler
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
