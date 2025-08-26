package main

import (
	_ "ai-chat-sql/internal/packed"

	_ "ai-chat-sql/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"

	"ai-chat-sql/internal/cmd"
)

func main() {
	gtime.SetTimeZone("Asia/Shanghai")
	cmd.Main.Run(gctx.GetInitCtx())
}
