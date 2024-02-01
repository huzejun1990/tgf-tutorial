// @Author huzejun 2024/1/31 15:43:00
package main

import (
	"github.com/huzejun1990/tgf/tgf-tutorial/common/services"
	"github.com/huzejun1990/tgf/tgf-tutorial/hall"
	"github.com/thkhxm/tgf/util"
)

func main() {
	//设置导出的golang目录
	util.SetAutoGenerateAPICodePath("../common/api")
	//设置导出的C#目录(用于Unity项目)
	//util.SetAutoGenerateAPICSCode("E:\\unity\\project\\t2\\Assets\\HotFix\\Code", "HotFix.Code")
	//根据接口生成对应的api结构
	util.GeneratorAPI[services.IHallService](hall.ModuleName, hall.Version, "api")

	//生成Cs代码(用于Unity)
	//util.GenerateCSApiService()
}
