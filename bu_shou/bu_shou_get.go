package xqkBuShou

import xqkAll "github.com/xqk/xqk-go-han-zi/xqk_all"

// GetByRune 根据rune获取部首
func GetByRune(r rune) string {
	return xqkAll.XqkBuShouGetByRune(r)
}
