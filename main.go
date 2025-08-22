package main

import (
	_ "ai-chat-sql/internal/packed"

	_ "ai-chat-sql/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"ai-chat-sql/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
