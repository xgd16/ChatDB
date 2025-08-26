package test

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/service"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
)

func TestPrompt_GetList(t *testing.T) {
	out, err := service.Prompt().GetList(consts.Ctx)
	if err != nil {
		t.Fatal(err)
	}
	g.DumpWithType(out)
}
