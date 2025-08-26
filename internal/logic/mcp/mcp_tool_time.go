package mcp

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *sMcpTool) GetNowTime(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error) {
	now := gtime.Now()

	out = mcp.NewToolResultText(gjson.MustEncodeString(g.Map{
		"datetime":        now.Format("Y-m-d H:i:s"),
		"UnixMillisecond": now.UnixMilli(),
	}))
	return
}
