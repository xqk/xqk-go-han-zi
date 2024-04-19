package xqkHanZiGen

import (
	xqkStrList "github.com/xqk/xqk-go-tool/string_list"
)

func GetMethodMap() map[string]map[string][]string {
	allMethodMap := make(map[string]map[string][]string)
	allMethodMap["xqkHanZi"] = map[string][]string{
		"Get": {
			"GetFirstLetterStr",
		},
		"Is": {
			"IsHanZi",
			"IsNotHanZi",
			"IsGt",
			"IsGe",
			"IsLt",
			"IsLe",
		},
		"To": {
			"ToPinYinList",
			"ToPinYinFirstLetterList",
			"ToBiHuaList",
			"ToBuShouList",
		},
	}
	allMethodMap["xqkHanZiList"] = map[string][]string{
		"Get": {
			"GetDesc",
			"GetAsc",
		},
		"Op": {
			"OpDesc",
			"OpAsc",
		},
	}
	allMethodMap["xqkBiHua"] = map[string][]string{
		"Get": {
			"GetByRune",
		},
	}
	allMethodMap["xqkBuShou"] = map[string][]string{
		"Get": {
			"GetByRune",
		},
	}
	allMethodMap["xqkPinYin"] = map[string][]string{
		"Get": {
			"GetByRune",
			"GetDefByRune",
			"GetFirstLetterByRune",
		},
		"Is": {
			"IsGt",
			"IsGe",
			"IsLt",
			"IsLe",
		},
		"To": {
			"ToNoTone",
			"ToLower",
			"ToUpper",
		},
	}
	// 名字排序
	for s1 := range allMethodMap {
		for s2 := range allMethodMap[s1] {
			allMethodMap[s1][s2] = xqkStrList.GetAscUnicode(allMethodMap[s1][s2])
		}
	}
	return allMethodMap
}
