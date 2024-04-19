package xqkHanZi

import xqkAll "github.com/xqk/xqk-go-han-zi/xqk_all"

// ToBiHuaList 汉字转笔画List，没有笔画的就为空
func ToBiHuaList(s string) []uint {
	return xqkAll.XqkHanZiToBiHuaList(s)
}

func ToBuShouList(s string) []string {
	return xqkAll.XqkHanZiToBuShouList(s)
}

// ToPinYinFirstLetterList 获取拼音首字母大写
func ToPinYinFirstLetterList(s string) []string {
	return xqkAll.XqkHanZiToPinYinFirstLetterList(s)
}

// ToPinYinList 汉字转拼音List，没有拼音的就为空
func ToPinYinList(s string) []string {
	return xqkAll.XqkHanZiToPinYinList(s)
}
