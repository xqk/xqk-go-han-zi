package xqkHanZiGen

import xqkGen "github.com/xqk/xqk-go-tool/xqk_gen"

func GetModuleList() []xqkGen.XqkModule {
	return []xqkGen.XqkModule{
		{
			PackageName: "xqkHanZi",
			DirName:     "han_zi",
			MethodModule: []xqkGen.XqkMethodModule{
				{
					Type:     xqkGen.Is,
					FileName: "han_zi_is.go",
				},
				{
					Type:     xqkGen.Get,
					FileName: "han_zi_get.go",
				},
				{
					Type:     xqkGen.To,
					FileName: "han_zi_to.go",
				},
			},
		},
		{
			PackageName: "xqkHanZiList",
			DirName:     "han_zi_list",
			MethodModule: []xqkGen.XqkMethodModule{
				{
					Type:     xqkGen.Get,
					FileName: "han_zi_list_get.go",
				},
				{
					Type:     xqkGen.Op,
					FileName: "han_zi_list_op.go",
				},
			},
		},
		{
			PackageName: "xqkPinYin",
			DirName:     "pin_yin",
			MethodModule: []xqkGen.XqkMethodModule{
				{
					Type:     xqkGen.Get,
					FileName: "pin_yin_get.go",
				},
				{
					Type:     xqkGen.Is,
					FileName: "pin_yin_is.go",
				},
				{
					Type:     xqkGen.To,
					FileName: "pin_yin_to.go",
				},
			},
		},
		{
			PackageName: "xqkBiHua",
			DirName:     "bi_hua",
			MethodModule: []xqkGen.XqkMethodModule{
				{
					Type:     xqkGen.Get,
					FileName: "bi_hua_get.go",
				},
			},
		},
		{
			PackageName: "xqkBuShou",
			DirName:     "bu_shou",
			MethodModule: []xqkGen.XqkMethodModule{
				{
					Type:     xqkGen.Get,
					FileName: "bu_shou_get.go",
				},
			},
		},
	}
}
