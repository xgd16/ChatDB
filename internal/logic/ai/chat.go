package ai

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/dao"
	"ai-chat-sql/internal/model"
	"ai-chat-sql/internal/service"
	"context"
	"time"

	"github.com/cloudwego/eino-ext/components/tool/mcp"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAiChat struct{}

func init() {
	service.RegisterAiChat(NewAiChat())
}

func NewAiChat() *sAiChat {
	return &sAiChat{}
}

// Chat 聊天
func (s *sAiChat) Chat(ctx context.Context, in model.ChatInput, respChan chan any) {
	// 创建响应通道
	HeartbeatCtx, cancel := context.WithCancel(ctx)
	s.AiChatHeartbeat(HeartbeatCtx, respChan)

	llm, err := service.AI().GetChatModel(in.Ai, in.Model)
	defer func() {
		if err != nil {
			respChan <- err
			close(respChan)
		}
	}()
	if err != nil {
		cancel()
		return
	}
	// 获取MCP工具
	mcpTools, err := mcp.GetTools(ctx, &mcp.Config{Cli: consts.McpClient})
	if err != nil {
		cancel()
		return
	}
	// 创建React智能体
	aiAgent, err := react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: llm,
		ToolsConfig:      compose.ToolsNodeConfig{Tools: mcpTools},
		MaxStep:          50,
	})
	if err != nil {
		cancel()
		return
	}
	// 获取需要操作的数据库信息
	dbTypeT, err := dao.DatabaseConf.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: 30 * time.Minute,
		Name:     "db_type:" + gconv.String(in.DatabaseId),
	}).Where("database_id = ?", in.DatabaseId).Value("db_type")
	if err != nil {
		cancel()
		return
	}

	prompt, err := service.Prompt().GetPrompt(ctx, consts.PromptMain)
	if err != nil {
		cancel()
		return
	}
	if g.IsEmpty(in.Prompt) {
		in.Prompt = "-"
	}
	out, err := aiAgent.Stream(ctx, []*schema.Message{
		{
			Role:    schema.System,
			Content: prompt.GetContent(in.DatabaseId, dbTypeT.String()),
		},
		{
			Role:    schema.System,
			Content: in.Prompt,
		},
		{
			Role:    schema.User,
			Content: in.Message,
		},
	})
	if err != nil {
		cancel()
		return
	}

	// AI输出流
	s.AiChatStreamOut(ctx, respChan, out, cancel)
}
