package xqkAll

import xqkStr "github.com/xqk/xqk-go-tool/string"

// XqkHanZiToPinYinList 汉字转拼音List，没有拼音的就为空
func XqkHanZiToPinYinList(s string) []string {
	var result []string
	rList := xqkStr.ToRuneList(s)
	for i := range rList {
		result = append(result, XqkPinYinGetDefByRune(rList[i]))
	}
	return result
}

// XqkHanZiToPinYinFirstLetterList 获取拼音首字母大写
func XqkHanZiToPinYinFirstLetterList(s string) []string {
	var result []string
	rList := xqkStr.ToRuneList(s)
	for i := range rList {
		result = append(result, XqkPinYinGetFirstLetterByRune(rList[i]))
	}
	return result
}

// XqkHanZiToBiHuaList 汉字转笔画List，没有笔画的就为空
func XqkHanZiToBiHuaList(s string) []uint {
	var result []uint
	rList := xqkStr.ToRuneList(s)
	for i := range rList {
		result = append(result, XqkBiHuaGetByRune(rList[i]))
	}
	return result
}

func XqkHanZiToBuShouList(s string) []string {
	var result []string
	rList := xqkStr.ToRuneList(s)
	for i := range rList {
		result = append(result, XqkBuShouGetByRune(rList[i]))
	}
	return result
}
