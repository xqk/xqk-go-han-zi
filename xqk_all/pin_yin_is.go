package xqkAll

import (
	"strings"

	xqkVar "github.com/xqk/xqk-go-han-zi/xqk_var"
	xqkRuneList "github.com/xqk/xqk-go-tool/rune_list"
	xqkStr "github.com/xqk/xqk-go-tool/string"
)

// XqkPinYinIsGt 判断拼音字符串是否大于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func XqkPinYinIsGt(s1, s2 string) bool {
	r1 := xqkStr.ToRuneList(s1)
	r2 := xqkStr.ToRuneList(s2)
	for i := range r1 {
		tr := xqkRuneList.GetElByIndex(r2, i)
		if r1[i] == tr {
			continue
		} else {
			// rune不一样了
			return pinYinRuneIsGt(r1[i], tr)
		}
	}
	return true
}

// XqkPinYinIsGe 判断拼音字符串是否大于等于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func XqkPinYinIsGe(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return XqkPinYinIsGt(s1, s2)
}

// XqkPinYinIsLt 判断拼音字符串是否小于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func XqkPinYinIsLt(s1, s2 string) bool {
	r1 := xqkStr.ToRuneList(s1)
	r2 := xqkStr.ToRuneList(s2)
	for i := range r1 {
		tr := xqkRuneList.GetElByIndex(r2, i)
		if r1[i] == tr {
			continue
		} else {
			// rune不一样了
			return !pinYinRuneIsGt(r1[i], tr)
		}
	}
	return false
}

// XqkPinYinIsLe 判断拼音字符串是否小于等于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func XqkPinYinIsLe(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return XqkPinYinIsLt(s1, s2)
}

func pinYinRuneIsGt(r1, r2 rune) bool {
	// 空数据要排在后面
	if r1 == 0 {
		return true
	}
	s1, s2 := string(r1), string(r2)
	p1 := xqkVar.PinYinToneMap[s1]
	p2 := xqkVar.PinYinToneMap[s2]
	if xqkStr.IsEmpty(p1) {
		p1 = s1
	}
	if xqkStr.IsEmpty(p2) {
		p2 = s2
	}
	if p1 == p2 {
		// 去掉音调后是一样的，那么就判断音调
		return xqkVar.PinYinToneIndexMap[s1] > xqkVar.PinYinToneIndexMap[s2]
	} else {
		// 去掉音调后不一样，那么就比较unicode码
		// 大小写一样
		l1 := strings.ToLower(p1)
		l2 := strings.ToLower(p2)
		if l1 == l2 {
			// 转成小写一样，如果1是大写（排在后面）就大于
			return xqkStr.IsUpperLetter(l1)
		} else {
			// 转成小写不一样，那么判断unicode
			return xqkRuneList.IsGtByUnicode(xqkStr.ToRuneList(l1), xqkStr.ToRuneList(l2))
		}
	}
}
