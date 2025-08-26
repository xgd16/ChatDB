package mcp

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
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

// TimestampToDateTime 时间戳转换为日期时间
func (s *sMcpTool) TimestampToDateTime(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error) {
	timestamp := request.GetString("timestamp", "")
	if timestamp == "" {
		return
	}
	gtime.NewFromTimeStamp(gconv.Int64(timestamp)).Format("Y-m-d H:i:s")
	out = mcp.NewToolResultText(gjson.MustEncodeString(g.Map{
		"datetime": gtime.NewFromTimeStamp(gconv.Int64(timestamp)).Format("Y-m-d H:i:s"),
	}))
	return
}
