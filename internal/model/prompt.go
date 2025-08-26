package model

import "fmt"

type PromptGetListOutputItem struct {
	FileName string `json:"fileName"`
	Path     string `json:"path"`
	Content  string `json:"content"`
}

func (s *PromptGetListOutputItem) SetParams(data ...any) *PromptGetListOutputItem {
	s.Content = fmt.Sprintf(s.Content, data...)
	return s
}
