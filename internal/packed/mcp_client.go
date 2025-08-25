package packed

import (
	"ai-chat-sql/internal/consts"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	g.Go(consts.Ctx, func(ctx context.Context) {
		var tryErr error
		for i := 0; i < 10; i++ {
			err := func() (err error) {
				mcpClient, err := client.NewSSEMCPClient(fmt.Sprintf("http://%s/sse", consts.Config.AiConfig.Mcp.Address))
				if err != nil {
					return
				}
				consts.McpClient = mcpClient
				if err = mcpClient.Start(consts.Ctx); err != nil {
					return
				}
				mcpClient.Initialize(consts.Ctx, mcp.InitializeRequest{})
				return
			}()
			if err != nil {
				tryErr = err
				time.Sleep(500 * time.Millisecond)
			} else {
				consts.Logger.Info(consts.Ctx, "初始化 MCP 客户端成功")
				break
			}
		}
		if tryErr != nil {
			consts.Logger.Errorf(consts.Ctx, "初始化 MCP 客户端失败 %s", tryErr.Error())
		}
	}, func(ctx context.Context, exception error) {
		consts.Logger.Errorf(consts.Ctx, "初始化 MCP 客户端失败 %s", exception.Error())
	})
}
