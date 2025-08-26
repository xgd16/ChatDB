package ai

import (
	"ai-chat-sql/internal/consts"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/cloudwego/eino/schema"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sAiChat) AiChatStreamOut(ctx context.Context, respChan chan any, stream *schema.StreamReader[*schema.Message]) {
	g.Go(ctx, func(ctx context.Context) {
		defer close(respChan)
		for {
			chunk, err := stream.Recv()
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
	}, func(ctx context.Context, exception error) {
		consts.Logger.Errorf(ctx, "AiChatStreamOut 异常 %s", exception.Error())
	})
}
