package xqkAll

import (
	"unicode"

	xqkVar "github.com/xqk/xqk-go-han-zi/xqk_var"
	xqkStr "github.com/xqk/xqk-go-tool/string"
)

// XqkPinYinToNoTone 将带音调的字符替换成没有音调的
func XqkPinYinToNoTone(str string) string {
	tpList := xqkStr.ToStrList(str)
	tp := ""
	for i := range tpList {
		if xqkVar.PinYinToneMap[tpList[i]] != "" {
			tp += xqkVar.PinYinToneMap[tpList[i]]
		} else {
			tp += tpList[i]
		}
	}
	return tp
}

// XqkPinYinToLower 将带音调的字符替换成没有音调的,并全换成小写
func XqkPinYinToLower(str string) string {
	str = XqkPinYinToNoTone(str)
	runeList := xqkStr.ToRuneList(str)
	for i := range runeList {
		runeList[i] = unicode.ToLower(runeList[i])
	}
	return string(runeList)
}

// XqkPinYinToUpper 将带音调的字符替换成没有音调的,并全换成大写
func XqkPinYinToUpper(str string) string {
	str = XqkPinYinToNoTone(str)
	runeList := xqkStr.ToRuneList(str)
	for i := range runeList {
		runeList[i] = unicode.ToUpper(runeList[i])
	}
	return string(runeList)
}
