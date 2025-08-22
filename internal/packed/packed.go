package packed

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/utility"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	// 将库文件构建出来
	if _, err := g.DB("master").Exec(consts.Ctx, "SELECT COUNT(*) AS table_count FROM sqlite_master WHERE type = 'table'"); err != nil {
		utility.PanicErr(err)
	}
}
