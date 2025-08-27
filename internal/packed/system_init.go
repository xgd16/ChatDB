package packed

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/service"
	"ai-chat-sql/utility"
)

func init() {
	if err := service.SystemInit().Init(consts.Ctx); err != nil {
		utility.PanicErr(err)
	}
}
