package ai

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/model"
	"ai-chat-sql/internal/service"
	"context"

	"github.com/cloudwego/eino-ext/components/tool/mcp"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/gogf/gf/v2/frame/g"
)

type sAiChat struct{}

func init() {
	service.RegisterAiChat(NewAiChat())
}

func NewAiChat() *sAiChat {
	return &sAiChat{}
}

// Chat 聊天
func (s *sAiChat) Chat(ctx context.Context, in model.ChatInput) (respChan chan any, err error) {
	llm, err := service.AI().GetChatModel(in.Ai, in.Model)
	if err != nil {
		return
	}
	// 获取MCP工具
	mcpTools, err := mcp.GetTools(ctx, &mcp.Config{Cli: consts.McpClient})
	if err != nil {
		return
	}
	// 创建React智能体
	aiAgent, err := react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: llm,
		ToolsConfig:      compose.ToolsNodeConfig{Tools: mcpTools},
	})
	if err != nil {
		return
	}

	prompt, err := service.Prompt().GetPrompt(ctx, consts.PromptMain)
	g.DumpWithType(prompt)
	if err != nil {
		return
	}
	if g.IsEmpty(in.Prompt) {
		in.Prompt = "-"
	}
	out, err := aiAgent.Stream(ctx, []*schema.Message{
		{
			Role:    schema.System,
			Content: prompt.GetContent(in.DatabaseId, "mysql"),
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
		return
	}

	// 创建响应通道
	respChan = make(chan any)
	// AI输出流
	s.AiChatStreamOut(ctx, respChan, out)

	return respChan, nil
}
