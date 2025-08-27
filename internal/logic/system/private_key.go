package system

import (
	"ai-chat-sql/internal/code"
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/service"
	"context"
)

type sPrivateKey struct {
}

func NewPrivateKey() *sPrivateKey {
	return &sPrivateKey{}
}

func init() {
	service.RegisterPrivateKey(NewPrivateKey())
}

// Verify 验证私钥
func (s *sPrivateKey) Verify(ctx context.Context, privateKey string) (err error) {
	if privateKey != consts.PrivateKey {
		return code.ToError(code.PrivateKeyError)
	}
	return
}
