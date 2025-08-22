package ai_chat

import (
	"ai-chat-sql/internal/consts"
	"ai-chat-sql/internal/dao"
	"ai-chat-sql/internal/model"
	"ai-chat-sql/internal/service"
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "ai-chat-sql/api/ai_chat/v1"
)

func (c *ControllerV1) SetDataBaseConfig(ctx context.Context, req *v1.SetDataBaseConfigReq) (res *v1.SetDataBaseConfigRes, err error) {
	res = &v1.SetDataBaseConfigRes{}
	// 检查数据库类型是否支持
	supportedTypes := map[string]bool{
		"mysql":    true,
		"postgres": true,
		"sqlite":   true,
	}

	if !supportedTypes[req.DbType] {
		return nil, gerror.Newf("不支持的数据库类型: %s, 支持的类型: mysql, postgres, sqlite", req.DbType)
	}

	// 检查是否已存在相同配置
	count, err := dao.DatabaseConf.Ctx(ctx).Count("host = ? AND port = ? AND db_name = ? AND db_type = ?", req.Host, req.Port, req.DbName, req.DbType)
	if err != nil {
		return
	}

	if count > 0 {
		err = gerror.New("已存在相同的数据库配置")
		return
	}

	// 创建新的数据库配置
	now := time.Now().Unix()

	// 使用map来避免零值问题，让数据库使用自增ID
	dbConfig := g.Map{
		"db_name":     req.DbName,
		"user_name":   req.UserName,
		"password":    req.Password,
		"host":        req.Host,
		"port":        req.Port,
		"db_type":     req.DbType,
		"create_time": int(now),
		"update_time": int(now),
	}

	// 保存到数据库
	if _, err = dao.DatabaseConf.Ctx(ctx).Data(dbConfig).Insert(); err != nil {
		return
	}

	return
}

func (c *ControllerV1) UpdateDataBaseConfig(ctx context.Context, req *v1.UpdateDataBaseConfigReq) (res *v1.UpdateDataBaseConfigRes, err error) {
	res = &v1.UpdateDataBaseConfigRes{}
	// 检查数据库类型是否支持
	supportedTypes := map[string]bool{
		"mysql":    true,
		"postgres": true,
		"sqlite":   true,
	}

	if !supportedTypes[req.DbType] {
		return nil, gerror.Newf("不支持的数据库类型: %s, 支持的类型: mysql, postgres, sqlite", req.DbType)
	}

	// 检查配置是否存在
	exists, err := dao.DatabaseConf.Ctx(ctx).Count("database_id = ?", req.DatabaseId)
	if err != nil {
		return
	}
	if exists == 0 {
		return nil, gerror.New("数据库配置不存在")
	}

	// 检查是否与其他配置冲突（排除当前配置）
	count, err := dao.DatabaseConf.Ctx(ctx).Count("host = ? AND port = ? AND db_name = ? AND db_type = ? AND database_id != ?",
		req.Host, req.Port, req.DbName, req.DbType, req.DatabaseId)
	if err != nil {
		return
	}

	if count > 0 {
		return nil, gerror.New("已存在相同的数据库配置")
	}

	// 更新数据库配置
	now := time.Now().Unix()
	updateData := g.Map{
		"db_name":     req.DbName,
		"user_name":   req.UserName,
		"password":    req.Password,
		"host":        req.Host,
		"port":        req.Port,
		"db_type":     req.DbType,
		"update_time": int(now),
	}

	if _, err = dao.DatabaseConf.Ctx(ctx).Where("database_id = ?", req.DatabaseId).Data(updateData).Update(); err != nil {
		return
	}

	return
}

func (c *ControllerV1) DeleteDataBaseConfig(ctx context.Context, req *v1.DeleteDataBaseConfigReq) (res *v1.DeleteDataBaseConfigRes, err error) {
	res = &v1.DeleteDataBaseConfigRes{}

	// 检查配置是否存在
	exists, err := dao.DatabaseConf.Ctx(ctx).Count("database_id = ?", req.DatabaseId)
	if err != nil {
		return
	}
	if exists == 0 {
		return nil, gerror.New("数据库配置不存在")
	}

	// 删除数据库配置
	if _, err = dao.DatabaseConf.Ctx(ctx).Where("database_id = ?", req.DatabaseId).Delete(); err != nil {
		return
	}

	return
}

func (c *ControllerV1) GetDataBaseConfigList(ctx context.Context, req *v1.GetDataBaseConfigListReq) (res *v1.GetDataBaseConfigListRes, err error) {
	res = &v1.GetDataBaseConfigListRes{}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	// 查询总数
	total, err := dao.DatabaseConf.Ctx(ctx).Count()
	if err != nil {
		return
	}

	// 查询分页数据
	list, err := dao.DatabaseConf.Ctx(ctx).
		Page(req.Page, req.Size).
		Order("database_id desc").
		All()
	if err != nil {
		return
	}

	// 转换数据结构
	var configs []*v1.SetDataBaseConfigReq
	if err = list.Structs(&configs); err != nil {
		return
	}

	res.List = configs
	res.Total = int(total)
	return
}

func (c *ControllerV1) GetAiModelList(ctx context.Context, req *v1.GetAiModelListReq) (res *v1.GetAiModelListRes, err error) {
	res = &v1.GetAiModelListRes{}

	t, err := consts.Cache.GetOrSetFunc(ctx, "open_ai_model_list", func(ctx context.Context) (result interface{}, err error) {
		result, err = service.AI().GetChatModeListJson(ctx, "openai")
		return
	}, time.Hour)
	if err != nil {
		return
	}
	data := struct {
		Data []*model.AIModelItem `json:"data"`
	}{}
	if err = json.Unmarshal(t.Bytes(), &data); err != nil {
		return
	}
	res.List = data.Data
	res.Total = len(res.List)
	return
}
