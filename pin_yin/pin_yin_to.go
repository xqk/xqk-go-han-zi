package xqkPinYin

import xqkAll "github.com/xqk/xqk-go-han-zi/xqk_all"

// ToLower 将带音调的字符替换成没有音调的,并全换成小写
func ToLower(str string) string {
	return xqkAll.XqkPinYinToLower(str)
}

// ToNoTone 将带音调的字符替换成没有音调的
func ToNoTone(str string) string {
	return xqkAll.XqkPinYinToNoTone(str)
}

// ToUpper 将带音调的字符替换成没有音调的,并全换成大写
func ToUpper(str string) string {
	return xqkAll.XqkPinYinToUpper(str)
}
