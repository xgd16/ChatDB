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
		// Md5Encode 对输入字符串计算 MD5（十六进制小写）
		Md5Encode(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// Base64Encode 对输入字符串进行 Base64 编码
		Base64Encode(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// Base64Decode 对输入 Base64 字符串进行解码
		Base64Decode(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// JwtParse 解析 JWT（不校验签名），输出 header/payload/rawHeader/rawPayload/rawSignature
		JwtParse(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// JsonEncode 将任意 JSON 解析为紧凑字符串输出（输入可以是 JSON 字符串）
		JsonEncode(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// ExecSql 执行SQL
		ExecSql(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// GetDatabaseInfo 获取数据库信息
		GetDatabaseInfo(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// ExportToExcel 将查询结果导出为 Excel 文件
		ExportToExcel(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// ExportData 通用导出接口（支持 XLSX 和 JSON）
		ExportData(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// ExecRedisCommand 执行Redis命令
		ExecRedisCommand(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// RunSafeShellCommand 执行安全受限的终端命令
		RunSafeShellCommand(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		GetNowTime(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// TimestampToDateTime 时间戳转换为日期时间
		TimestampToDateTime(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
		// GetCalendarDays 获取指定年月的每一天日期
		GetCalendarDays(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error)
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
