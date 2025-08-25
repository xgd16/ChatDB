package mcp

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/model"
	"ai-chat-sql/internal/service"
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type sMcpHandler struct {
}

func NewMcpHandler() *sMcpHandler {
	return &sMcpHandler{}
}

func init() {
	service.RegisterMcpHandler(NewMcpHandler())
	service.RegisterMcpTool(NewMcpTool())
}

func (s *sMcpHandler) GetList() []model.McpReg {
	return []model.McpReg{
		{
			Name:        "SQL_Actuator",
			Description: "Convert the user's requirements into SQL statements, execute the SQL statements, and return the execution results",
			ToolOptions: []mcp.ToolOption{
				mcp.WithString("databaseId",
					mcp.Required(),
					mcp.Description("The ID of the database to operate on (databaseId)"),
				),
				mcp.WithString("sql",
					mcp.Required(),
					mcp.Description("The SQL statement to be executed"),
				),
			},
			Fn: service.McpTool().ExecSql,
		},
		{
			Name:        "hello_world",
			Description: "Say hello to someone",
			ToolOptions: []mcp.ToolOption{
				mcp.WithString("name",
					mcp.Required(),
					mcp.Description("Name of the person to greet"),
				),
			},
			Fn: s.SayHello,
		},
	}
}

func (s *sMcpHandler) GetMcpFn(item *model.McpReg) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (result *mcp.CallToolResult, err error) {
		defer func() {
			if err := recover(); err != nil {
				consts.Logger.Printf(ctx, "panic error %+v", err)
			}
		}()
		consts.Logger.Printf(ctx, "接收到请求 %s 请求内容 %+v", item.Name, request.Params.Arguments)
		if item.Fn == nil {
			return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
				return mcp.NewToolResultText("处理函数未定义"), nil
			}(ctx, request)
		}
		return item.Fn(ctx, request)
	}
}

func (s *sMcpHandler) SayHello(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name := request.GetString("name", "")
	if name == "" {
		return nil, errors.New("name must be a string")
	}
	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}

type sMcpTool struct{}

func NewMcpTool() *sMcpTool {
	return &sMcpTool{}
}
