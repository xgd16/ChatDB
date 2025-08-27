package code

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var baseStartCode = 9999    // 让code从 10000 开始
var loggerStartCode = 19999 // 让code从 20000 开始
var RespCodeList = gmap.NewListMap(true)

func newCode(message string, ptr *int) gcode.Code {
	code := getNewCode(ptr)
	RespCodeList.Set(code, message)
	return gcode.New(code, message, nil)
}

func getNewCode(code *int) int {
	*code += 1
	return *code
}

func ToError(code gcode.Code, text ...string) error {
	return gerror.NewCode(code, text...)
}
