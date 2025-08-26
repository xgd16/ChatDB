package model

import "fmt"

type PromptGetListOutputItem struct {
	FileName string `json:"fileName"`
	Path     string `json:"path"`
	Content  string `json:"content"`
}

func (s *PromptGetListOutputItem) GetContent(data ...any) string {
	return fmt.Sprintf(s.Content, data...)
}
