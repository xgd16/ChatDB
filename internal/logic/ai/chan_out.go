package ai

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/model"
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func (s *sAiChat) AiChatStreamOut(ctx context.Context, respChan chan any, stream *schema.StreamReader[*schema.Message], cancel context.CancelFunc) {
	g.Go(ctx, func(ctx context.Context) {
		defer close(respChan)
		for {
			chunk, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				// 发送结束包
				if err = model.SendChatOutDataItem("end", "", "", respChan); err != nil {
					return
				}
				cancel()
				break
			}
			if err != nil {
				fmt.Printf("接收数据错误: %v\n", err)
				cancel()
				return
			}

			// 发送chunk到通道
			select {
			case respChan <- model.ChatOutDataItem{
				Event:      "message",
				Role:       gconv.String(chunk.Role),
				Content:    chunk.Content,
				CreateTime: gtime.TimestampMilli(),
			}:
			case <-ctx.Done():
				close(respChan)
				return
			}
		}
	}, func(ctx context.Context, exception error) {
		cancel()
		consts.Logger.Errorf(ctx, "AiChatStreamOut 异常 %s", exception.Error())
	})
}

func (s *sAiChat) AiChatHeartbeat(ctx context.Context, respChan chan any) {
	g.Go(ctx, func(ctx context.Context) {
		ticker := time.NewTicker(time.Millisecond * 1500)
		respChan <- "event: ping"
		for {
			select {
			case <-ticker.C:
				respChan <- "event: ping"
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}, func(ctx context.Context, exception error) {
		consts.Logger.Errorf(ctx, "AiChatHeartbeat 异常 %s", exception.Error())
	})
}
