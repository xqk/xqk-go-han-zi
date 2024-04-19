package xqkHanZiList

import xqkAll "github.com/xqk/xqk-go-han-zi/xqk_all"

// OpAsc 将字符串数组按拼音升序排序
func OpAsc(list *[]string) {
	xqkAll.XqkHanZiListOpAsc(list)
}

// OpDesc 将字符串数组按拼音降序排序
func OpDesc(list *[]string) {
	xqkAll.XqkHanZiListOpDesc(list)
}
