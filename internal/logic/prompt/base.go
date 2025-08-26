package prompt

import (
	"ai-chat-sql/internal/model"
	"ai-chat-sql/internal/service"
	"context"
	"errors"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

type sPrompt struct {
	promptList []*model.PromptGetListOutputItem
	promptMap  map[string]*model.PromptGetListOutputItem
	mutex      *sync.RWMutex
	promptDir  string
}

func NewPrompt() *sPrompt {
	return &sPrompt{
		promptList: make([]*model.PromptGetListOutputItem, 0),
		promptMap:  make(map[string]*model.PromptGetListOutputItem),
		mutex:      new(sync.RWMutex),
		promptDir:  "./prompt",
	}
}

func init() {
	service.RegisterPrompt(NewPrompt())
}

// GetList 获取提示词列表
func (s *sPrompt) GetList(ctx context.Context) (out []*model.PromptGetListOutputItem, err error) {
	s.mutex.RLock()
	if len(s.promptList) > 0 {
		out = s.promptList
		s.mutex.RUnlock()
		return
	}
	s.mutex.RUnlock()
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// 读取文件夹
	filePaths, err := gfile.ScanDirFile(s.promptDir, "*.md")
	if err != nil {
		return
	}
	out = make([]*model.PromptGetListOutputItem, 0)
	for _, v := range filePaths {
		fileName := gfile.Name(v)
		content := gfile.GetContents(v)
		out = append(out, &model.PromptGetListOutputItem{
			FileName: fileName,
			Path:     v,
			Content:  content,
		})
	}
	s.promptList = out
	return
}

// CleanCache 清理缓存
func (s *sPrompt) CleanCache() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.promptList = make([]*model.PromptGetListOutputItem, 0)
	s.promptMap = make(map[string]*model.PromptGetListOutputItem)
}

// GetPrompt 获取提示词
func (s *sPrompt) GetPrompt(ctx context.Context, fileName string) (out *model.PromptGetListOutputItem, err error) {
	// 从缓存中获取
	var ok bool
	if !g.IsEmpty(s.promptMap) {
		s.mutex.RLock()
		defer s.mutex.RUnlock()
		out, ok = s.promptMap[fileName]
		if !ok {
			err = errors.New("prompt not found")
			return
		}
		return
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, v := range s.promptList {
		s.promptMap[v.FileName] = v
	}
	out, ok = s.promptMap[fileName]
	if !ok {
		err = errors.New("prompt not found")
		return
	}
	return
}
