package system

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type sSystemInit struct {
}

func NewSystemInit() *sSystemInit {
	return &sSystemInit{}
}

func init() {
	service.RegisterSystemInit(NewSystemInit())
}

// Init 初始化
func (s *sSystemInit) Init(ctx context.Context) (err error) {
	if err = s.CreatePrivateKey(ctx); err != nil {
		return
	}
	return
}

// CreatePrivateKey 创建私钥
func (s *sSystemInit) CreatePrivateKey(ctx context.Context) (err error) {
	const privateKeyPath = "./private.key"
	if !gfile.Exists(privateKeyPath) {
		consts.PrivateKey = gbase64.EncodeString(fmt.Sprintf(
			"%s-%s",
			grand.S(64, true), gconv.String(gtime.Now().UnixMilli()),
		))
		if err = gfile.PutContents(privateKeyPath, consts.PrivateKey); err != nil {
			return
		}
		return
	}
	consts.PrivateKey = gfile.GetContents(privateKeyPath)
	return
}
