package xqkAll

import (
	"sort"

	xqkStrList "github.com/xqk/xqk-go-tool/string_list"
)

// XqkHanZiListGetDesc 按照拼音将字符串降序排序
func XqkHanZiListGetDesc(list []string) []string {
	if xqkStrList.IsEmpty(list) {
		return nil
	}
	sort.Slice(list, func(i, j int) bool {
		return XqkHanZiIsGe(list[i], list[j])
	})
	return list
}

// XqkHanZiListGetAsc 按照拼音将字符串升序排序
func XqkHanZiListGetAsc(list []string) []string {
	if xqkStrList.IsEmpty(list) {
		return nil
	}
	sort.Slice(list, func(i, j int) bool {
		return XqkHanZiIsLe(list[i], list[j])
	})
	return list
}
