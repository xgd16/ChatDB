package mcp

import (
	"ai-chat-sql/internal/consts"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/xuri/excelize/v2"
)

// ExportToExcel 将查询结果导出为 Excel 文件
func (s *sMcpTool) ExportToExcel(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error) {
	// 获取导出参数
	dataJson := request.GetString("data", "")
	if dataJson == "" {
		err = errors.New("data parameter is required")
		return
	}

	fileName := request.GetString("fileName", "export_"+time.Now().Format("20060102_150405"))
	if fileName == "" {
		fileName = "export_" + time.Now().Format("20060102_150405")
	}

	// 解析数据
	var data []map[string]interface{}
	if err = gjson.Unmarshal([]byte(dataJson), &data); err != nil {
		err = errors.New("invalid data format: " + err.Error())
		return
	}

	if len(data) == 0 {
		err = errors.New("data is empty")
		return
	}

	// 创建 Excel 文件
	f := excelize.NewFile()
	defer func() {
		if err = f.Close(); err != nil {
			consts.Logger.Error(ctx, "close excel file error: "+err.Error())
		}
	}()

	// 获取第一个 Sheet（默认为 Sheet1）
	sheet := "Sheet1"
	idx, err := f.GetSheetIndex(sheet)
	if err == nil && idx >= 0 {
		f.SetActiveSheet(idx)
	}

	// 获取所有列名（从第一行数据的键）
	var columns []string
	for key := range data[0] {
		columns = append(columns, key)
	}

	// 写入表头
	for i, col := range columns {
		colName, _ := excelize.ColumnNumberToName(i + 1)
		cell := colName + "1"
		f.SetCellValue(sheet, cell, col)
		// 设置表头样式
		style, _ := f.NewStyle(&excelize.Style{
			Font:      &excelize.Font{Bold: true, Color: "FFFFFF"},
			Fill:      excelize.Fill{Type: "pattern", Color: []string{"4472C4"}, Pattern: 1},
			Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
			Border: []excelize.Border{
				{Type: "left", Color: "000000", Style: 1},
				{Type: "right", Color: "000000", Style: 1},
				{Type: "top", Color: "000000", Style: 1},
				{Type: "bottom", Color: "000000", Style: 1},
			},
		})
		f.SetCellStyle(sheet, cell, cell, style)
	}

	// 写入数据
	bodyStyle, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "D3D3D3", Style: 1},
			{Type: "right", Color: "D3D3D3", Style: 1},
			{Type: "top", Color: "D3D3D3", Style: 1},
			{Type: "bottom", Color: "D3D3D3", Style: 1},
		},
	})

	for rowIdx, row := range data {
		for colIdx, col := range columns {
			colName, _ := excelize.ColumnNumberToName(colIdx + 1)
			cell := fmt.Sprintf("%s%d", colName, rowIdx+2)
			value := row[col]
			f.SetCellValue(sheet, cell, value)
			f.SetCellStyle(sheet, cell, cell, bodyStyle)
		}
	}

	// 设置列宽
	for i := range columns {
		col, _ := excelize.ColumnNumberToName(i + 1)
		f.SetColWidth(sheet, col, col, 20)
	}

	// 设置行高
	f.SetRowHeight(sheet, 1, 25)

	// 生成文件路径
	exportDir := "exports"
	if err = os.MkdirAll(exportDir, os.ModePerm); err != nil {
		err = errors.New("create export directory error: " + err.Error())
		return
	}

	filePath := filepath.Join(exportDir, fileName+".xlsx")

	// 保存文件
	if err = f.SaveAs(filePath); err != nil {
		err = errors.New("save excel file error: " + err.Error())
		return
	}

	// 返回文件路径
	result := g.Map{
		"fileName":    fileName + ".xlsx",
		"filePath":    filePath,
		"downloadUrl": "/api/ai_chat/v1/export/download?file=" + fileName + ".xlsx",
		"message":     fmt.Sprintf("导出成功！共 %d 行数据", len(data)),
	}

	out = mcp.NewToolResultText(gjson.MustEncodeString(result))
	consts.Logger.Info(ctx, fmt.Sprintf("导出 Excel 文件: %s", filePath))
	return
}

// ExportData 通用导出接口（支持 XLSX 和 JSON）
func (s *sMcpTool) ExportData(ctx context.Context, request mcp.CallToolRequest) (out *mcp.CallToolResult, err error) {
	dataJson := request.GetString("data", "")
	fileName := request.GetString("fileName", "export_"+time.Now().Format("20060102_150405"))
	exportFormat := request.GetString("format", "xlsx") // xlsx 或 json

	if dataJson == "" {
		err = errors.New("data parameter is required")
		return
	}

	// 生成文件路径
	exportDir := "exports"
	if err = os.MkdirAll(exportDir, os.ModePerm); err != nil {
		err = errors.New("create export directory error: " + err.Error())
		return
	}

	var fileExt string
	var filePath string

	if exportFormat == "json" {
		fileExt = ".json"
		filePath = filepath.Join(exportDir, fileName+fileExt)

		// 写入 JSON 文件
		if err = os.WriteFile(filePath, []byte(dataJson), 0644); err != nil {
			err = errors.New("write json file error: " + err.Error())
			return
		}
	} else {
		// 默认使用 XLSX
		fileExt = ".xlsx"
		filePath = filepath.Join(exportDir, fileName+fileExt)

		// 解析数据并写入 XLSX
		var data []map[string]interface{}
		if err = gjson.Unmarshal([]byte(dataJson), &data); err != nil {
			err = errors.New("invalid data format: " + err.Error())
			return
		}

		if len(data) == 0 {
			err = errors.New("data is empty")
			return
		}

		f := excelize.NewFile()
		defer func() {
			if err = f.Close(); err != nil {
				consts.Logger.Error(ctx, "close excel file error: "+err.Error())
			}
		}()

		sheet := "Sheet1"

		// 获取列名
		var columns []string
		for key := range data[0] {
			columns = append(columns, key)
		}

		// 写入表头
		for i, col := range columns {
			colName, _ := excelize.ColumnNumberToName(i + 1)
			cell := colName + "1"
			f.SetCellValue(sheet, cell, col)
			style, _ := f.NewStyle(&excelize.Style{
				Font:      &excelize.Font{Bold: true, Color: "FFFFFF"},
				Fill:      excelize.Fill{Type: "pattern", Color: []string{"4472C4"}, Pattern: 1},
				Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
			})
			f.SetCellStyle(sheet, cell, cell, style)
		}

		// 写入数据
		for rowIdx, row := range data {
			for colIdx, col := range columns {
				colName, _ := excelize.ColumnNumberToName(colIdx + 1)
				cell := fmt.Sprintf("%s%d", colName, rowIdx+2)
				f.SetCellValue(sheet, cell, row[col])
			}
		}

		// 设置列宽
		for i := range columns {
			col, _ := excelize.ColumnNumberToName(i + 1)
			f.SetColWidth(sheet, col, col, 20)
		}

		// 保存文件
		if err = f.SaveAs(filePath); err != nil {
			err = errors.New("save excel file error: " + err.Error())
			return
		}
	}

	// 返回下载链接
	result := g.Map{
		"fileName":    fileName + fileExt,
		"filePath":    filePath,
		"downloadUrl": "/api/ai_chat/v1/export/download?file=" + fileName + fileExt,
		"message":     "导出成功！",
	}

	out = mcp.NewToolResultText(gjson.MustEncodeString(result))
	consts.Logger.Info(ctx, fmt.Sprintf("导出文件: %s", filePath))
	return
}
