package ai

import (
	"ai-chat-sql/internal/model"
	"ai-chat-sql/internal/service"
	"context"
	"errors"
	"fmt"
	"io"

	einoModel "github.com/cloudwego/eino/components/model"
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
	chatModel, err := service.AI().GetChatModel(in.Ai)
	if err != nil {
		return
	}
	out, err := chatModel.Stream(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: in.Prompt,
		},
	}, einoModel.WithModel(in.Model))
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
