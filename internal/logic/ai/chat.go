package ai

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/model"
	"ai-chat-sql/internal/service"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/cloudwego/eino-ext/components/tool/mcp"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
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

	out, err := aiAgent.Stream(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: in.Prompt,
		},
	})
	if err != nil {
		return
	}

	// 创建响应通道
	respChan = make(chan any)

	// 启动goroutine处理输出
	go func() {
		defer close(respChan)

		for {
			chunk, err := out.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				fmt.Printf("接收数据错误: %v\n", err)
				return
			}

			// 发送chunk到通道
			select {
			case respChan <- chunk:
			case <-ctx.Done():
				return
			}
		}
	}()

	return respChan, nil
}
