package xqkAll

import xqkVar "github.com/xqk/xqk-go-han-zi/xqk_var"

func XqkBiHuaGetByRune(r rune) uint {
	if XqkHanZiIsNotHanZi(r) {
		return 0
	}
	return xqkVar.FuncHanZiGetBiHua(r)
}
