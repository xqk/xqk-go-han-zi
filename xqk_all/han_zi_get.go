package xqkAll

import xqkStrList "github.com/xqk/xqk-go-tool/string_list"

// XqkHanZiGetFirstLetterStr 获取首字母大写字符串
func XqkHanZiGetFirstLetterStr(s string) string {
	return xqkStrList.ToStr(xqkStrList.GetDeleteEmptyEl(XqkHanZiToPinYinFirstLetterList(s)))
}
