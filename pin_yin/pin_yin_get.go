package xqkPinYin

import xqkAll "github.com/xqk/xqk-go-han-zi/xqk_all"

// GetByRune 根据rune获取拼音
func GetByRune(r rune) []string {
	return xqkAll.XqkPinYinGetByRune(r)
}

// GetDefByRune 根据rune获取默认拼音
func GetDefByRune(r rune) string {
	return xqkAll.XqkPinYinGetDefByRune(r)
}

// GetFirstLetterByRune 根据rune获取拼音的首字母大写
func GetFirstLetterByRune(r rune) string {
	return xqkAll.XqkPinYinGetFirstLetterByRune(r)
}
