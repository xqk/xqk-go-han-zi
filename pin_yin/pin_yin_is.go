package xqkPinYin

import xqkAll "github.com/xqk/xqk-go-han-zi/xqk_all"

// IsGe 判断拼音字符串是否大于等于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func IsGe(s1, s2 string) bool {
	return xqkAll.XqkPinYinIsGe(s1, s2)
}

// IsGt 判断拼音字符串是否大于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func IsGt(s1, s2 string) bool {
	return xqkAll.XqkPinYinIsGt(s1, s2)
}

// IsLe 判断拼音字符串是否小于等于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func IsLe(s1, s2 string) bool {
	return xqkAll.XqkPinYinIsLe(s1, s2)
}

// IsLt 判断拼音字符串是否小于
// - 大小写一样，小写优于大写
// - 声调按顺序排序
func IsLt(s1, s2 string) bool {
	return xqkAll.XqkPinYinIsLt(s1, s2)
}
