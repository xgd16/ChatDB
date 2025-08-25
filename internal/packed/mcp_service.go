package packed

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func init() {
	g.Go(consts.Ctx, func(ctx context.Context) {
		// 启动MCP服务
		// Create MCP server
		s := server.NewMCPServer(
			"MCP Server 🚀",
			"1.0.0",
		)

		// Add tool
		fmt.Printf("–––––––––––––––––––––––––––––––––MCP SERVER–––––––––––––––––––––––––––––––––\n\n")
		for _, item := range service.McpHandler().GetList() {
			fmt.Printf("添加工具 %s - %s\n", item.Name, item.Description)
			s.AddTool(mcp.NewTool(item.Name,
				append([]mcp.ToolOption{
					mcp.WithDescription(item.Description),
				}, item.ToolOptions...)...,
			), service.McpHandler().GetMcpFn(&item))
		}
		fmt.Printf("\n––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––\n")

		fmt.Printf("MCP SSE服务已启动地址: http://%s/sse\n", consts.Config.AiConfig.Mcp.Address)

		// Start the stdio server
		if err := server.NewSSEServer(s).Start(consts.Config.AiConfig.Mcp.Address); err != nil {
			panic(err)
		}
	}, func(ctx context.Context, exception error) {
		consts.Logger.Errorf(ctx, "初始化 MCP 服务失败 %s", exception.Error())
	})
}
