package test

import (
	"ai-chat-sql/internal/consts"
	_ "ai-chat-sql/internal/logic"
	"ai-chat-sql/internal/service"
	"testing"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestAi_OpenAI(t *testing.T) {
	chatModel, err := service.AI().GetChatModel("openai", "")
	gtest.AssertNil(err)
	g.DumpWithType(chatModel.Generate(consts.Ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "你好",
		},
	}, model.WithModel("gpt-5-mini")))
}

func TestAi_DeepSeek(t *testing.T) {
	chatModel, err := service.AI().GetChatModel("deepseek", "")
	gtest.AssertNil(err)
	g.DumpWithType(chatModel.Generate(consts.Ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "你好",
		},
	}, model.WithModel("deepseek-chat")))
}

func TestAi_GetChatModeListJson(t *testing.T) {
	json, err := service.AI().GetChatModeListJson(consts.Ctx, "openai")
	gtest.AssertNil(err)
	g.DumpWithType(json)
}
