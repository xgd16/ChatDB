package ai

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/service"
	"context"
	"errors"
	"fmt"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino-ext/components/model/openai"
	einoModel "github.com/cloudwego/eino/components/model"
	"github.com/gogf/gf/v2/net/gclient"
)

type sAI struct {
}

func init() {
	service.RegisterAI(NewAI())
}

func NewAI() *sAI {
	return &sAI{}
}

// GetChatModel 获取聊天模型
func (s *sAI) GetChatModel(ai string) (chatModel einoModel.BaseChatModel, err error) {
	switch ai {
	case "openai":
		chatModel, err = openai.NewChatModel(consts.Ctx, &openai.ChatModelConfig{
			BaseURL: consts.Config.AiConfig.OpenAI.BaseUrl,
			Model:   "gpt-4o-mini",
			APIKey:  consts.Config.AiConfig.OpenAI.Key,
		})
		return
	case "deepseek":
		chatModel, err = deepseek.NewChatModel(consts.Ctx, &deepseek.ChatModelConfig{
			BaseURL: consts.Config.AiConfig.DeepSeek.BaseUrl,
			Model:   "deepseek-chat",
			APIKey:  consts.Config.AiConfig.DeepSeek.Key,
		})
		return
	}
	return
}

// GetChatModeListJson 获取聊天模型列表
func (s *sAI) GetChatModeListJson(ctx context.Context, ai string) (json string, err error) {
	var (
		url string
		key string
	)
	switch ai {
	case "openai":
		url = consts.Config.AiConfig.OpenAI.BaseUrl
		key = consts.Config.AiConfig.OpenAI.Key
	case "deepseek":
		url = consts.Config.AiConfig.DeepSeek.BaseUrl
		key = consts.Config.AiConfig.DeepSeek.Key
	default:
		err = errors.New("未知的操作")
		return
	}
	resp, err := gclient.New().Discovery(nil).SetHeader("authorization", key).Get(ctx, fmt.Sprintf("%s/models", url))
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("获取模型列表失败")
		return
	}
	json = resp.ReadAllString()
	return
}
