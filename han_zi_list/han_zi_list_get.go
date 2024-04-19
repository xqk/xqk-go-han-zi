package xqkHanZiList

import xqkAll "github.com/xqk/xqk-go-han-zi/xqk_all"

// GetAsc 按照拼音将字符串升序排序
func GetAsc(list []string) []string {
	return xqkAll.XqkHanZiListGetAsc(list)
}

// GetDesc 按照拼音将字符串降序排序
func GetDesc(list []string) []string {
	return xqkAll.XqkHanZiListGetDesc(list)
}
