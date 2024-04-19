package xqkAll

import (
	"strings"

	xqkVar "github.com/xqk/xqk-go-han-zi/xqk_var"
	xqkStr "github.com/xqk/xqk-go-tool/string"
)

// XqkPinYinGetByRune 根据rune获取拼音
func XqkPinYinGetByRune(r rune) []string {
	if XqkHanZiIsNotHanZi(r) {
		return nil
	}
	return xqkVar.FuncHanZiGetPinYin(r)
}

// XqkPinYinGetDefByRune 根据rune获取默认拼音
func XqkPinYinGetDefByRune(r rune) string {
	if XqkHanZiIsNotHanZi(r) {
		return ""
	}
	pList := xqkVar.FuncHanZiGetPinYin(r)
	if len(pList) == 0 {
		return ""
	} else {
		return pList[0]
	}
}

// XqkPinYinGetFirstLetterByRune 根据rune获取拼音的首字母大写
func XqkPinYinGetFirstLetterByRune(r rune) string {
	ts := xqkStr.GetFirstRuneStr(XqkPinYinToNoTone(XqkPinYinGetDefByRune(r)))
	if xqkStr.IsLowerLetter(ts) {
		return strings.ToUpper(ts)
	}
	if xqkStr.IsUpperLetter(ts) {
		return ts
	}
	return ""
}
