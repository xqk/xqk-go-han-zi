package xqkAll

import (
	xqkRuneList "github.com/xqk/xqk-go-tool/rune_list"
	xqkStr "github.com/xqk/xqk-go-tool/string"
)

// XqkHanZiIsHanZi 是否是汉字
func XqkHanZiIsHanZi(i rune) bool {
	// 统一A
	if 13312 <= i && i <= 19903 {
		return true
	}
	// 统一
	if 19968 <= i && i <= 40959 {
		return true
	}
	// 兼容
	if 63744 <= i && i <= 64255 {
		return true
	}
	// 统一B
	if 131072 <= i && i <= 173791 {
		return true
	}
	// 统一C
	if 173824 <= i && i <= 177983 {
		return true
	}
	// 统一D
	if 177984 <= i && i <= 178207 {
		return true
	}
	// 统一E
	if 178208 <= i && i <= 183983 {
		return true
	}
	// 统一F
	if 183984 <= i && i <= 191471 {
		return true
	}
	// 兼容增补
	if 194560 <= i && i <= 195103 {
		return true
	}
	// 统一G
	if 196608 <= i && i <= 201551 {
		return true
	}
	return false
}

// XqkHanZiIsNotHanZi 是否不是汉字
func XqkHanZiIsNotHanZi(i rune) bool {
	return !XqkHanZiIsHanZi(i)
}

// XqkHanZiIsGt 字符串是否大于
//
// - 汉字在前，非汉字在后
// - 按拼音、音调、部首、笔画排序
func XqkHanZiIsGt(s1, s2 string) bool {
	if xqkStr.IsBlank(s1) {
		return true
	}
	r1 := xqkStr.ToRuneList(s1)
	r2 := xqkStr.ToRuneList(s2)
	for i := range r1 {
		// 不是汉字排到后面去
		if XqkHanZiIsNotHanZi(r1[i]) {
			return true
		}
		tr := xqkRuneList.GetElByIndex(r2, i)
		if r1[i] == tr {
			continue
		} else {
			// rune不一样了
			// 比较拼音
			p1 := XqkPinYinGetDefByRune(r1[i])
			p2 := XqkPinYinGetDefByRune(tr)
			if p1 == p2 {
				// 如果拼音相等，比较部首
				b1 := XqkBuShouGetByRune(r1[i])
				b2 := XqkBuShouGetByRune(tr)
				if b1 == b2 {
					// 如果部首相等，比较笔画
					h1 := XqkBiHuaGetByRune(r1[i])
					h2 := XqkBiHuaGetByRune(tr)
					return h1 > h2
				} else {
					return xqkStr.IsGt(b1, b2)
				}
			} else {
				return XqkPinYinIsGt(p1, p2)
			}
		}
	}
	return true
}

// XqkHanZiIsGe 字符串是否大于等于
//
// - 汉字在前，非汉字在后
// - 按拼音、音调、部首、笔画排序
func XqkHanZiIsGe(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return XqkHanZiIsGt(s1, s2)
}

// XqkHanZiIsLt 字符串是否小于
//
// - 汉字在前，非汉字在后
// - 按拼音、音调、部首、笔画排序
func XqkHanZiIsLt(s1, s2 string) bool {
	return !XqkHanZiIsGt(s1, s2)
}

// XqkHanZiIsLe 字符串是否小于等于
//
// - 汉字在前，非汉字在后
// - 按拼音、音调、部首、笔画排序
func XqkHanZiIsLe(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return !XqkHanZiIsLt(s1, s2)
}
