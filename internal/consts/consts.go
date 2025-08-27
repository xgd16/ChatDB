package consts

import (
	"ai-chat-sql/internal/model"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	mcpclient "github.com/mark3labs/mcp-go/client"
)

var (
	Ctx          = gctx.New()
	Logger       = g.Log()
	SystemConfig *gjson.Json
	Config       *model.ConfigData
)

var (
	McpClient *mcpclient.Client
)

var (
	PrivateKey string = ""
)

const JwtSubjectUser = "ai-chat-user"

func init() {
	if err := initSystemConfig(); err != nil {
		panic(err)
	}
}

func initSystemConfig() (err error) {
	data, err := g.Cfg().Data(Ctx)
	if err != nil {
		return
	}
	SystemConfig = gjson.New(data)
	if err = SystemConfig.Scan(&Config); err != nil {
		return
	}
	return
}
