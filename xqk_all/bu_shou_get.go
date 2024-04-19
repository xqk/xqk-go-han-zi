package xqkAll

import xqkVar "github.com/xqk/xqk-go-han-zi/xqk_var"

// XqkBuShouGetByRune 根据rune获取部首
func XqkBuShouGetByRune(r rune) string {
	if XqkHanZiIsNotHanZi(r) {
		return ""
	}
	return xqkVar.FuncHanZiGetBuShou(r)
}
