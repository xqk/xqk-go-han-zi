package xqkAll

// XqkHanZiListOpDesc 将字符串数组按拼音降序排序
func XqkHanZiListOpDesc(list *[]string) {
	if list == nil {
		return
	}
	*list = XqkHanZiListGetDesc(*list)
}

// XqkHanZiListOpAsc 将字符串数组按拼音升序排序
func XqkHanZiListOpAsc(list *[]string) {
	if list == nil {
		return
	}
	*list = XqkHanZiListGetAsc(*list)
}
