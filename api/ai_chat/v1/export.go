package v1

import "github.com/gogf/gf/v2/frame/g"

// ExportDownloadReq 下载导出文件请求
type ExportDownloadReq struct {
	g.Meta `path:"/export/download" method:"get" tags:"导出"`
	File   string `json:"file" v:"required" dc:"导出文件名"`
}

// ExportDownloadRes 下载导出文件响应
type ExportDownloadRes struct{}
