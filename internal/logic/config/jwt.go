package config

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/model"
	"context"
	"errors"
)

// GetJwtOptions 获取JWT配置
func (s *sConfig) GetJwtOptions(ctx context.Context, secret string) (option *model.JwtOption, err error) {
	v, err := consts.Cache.GetOrSetFunc(ctx, "jwt_cfg_options", func(ctx context.Context) (value interface{}, err error) {
		newMap := make(map[string]*model.JwtOption)
		for _, option := range consts.Config.Jwt {
			newMap[option.Subject] = option
		}
		value = newMap
		return
	}, 0)
	m := v.Interface().(map[string]*model.JwtOption)
	option, ok := m[secret]
	if !ok {
		err = errors.New("jwt option not found")
		return
	}
	return
}
