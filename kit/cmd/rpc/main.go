// @Author huzejun 2024/1/31 23:43:00
package main

import (
	"github.com/huzejun1990/tgf/tgf-tutorial/common/services"
	"github.com/huzejun1990/tgf/tgf-tutorial/prop"
	"github.com/thkhxm/tgf/util"
)

func main() {
	//设置导出的golang目录
	util.SetAutoGenerateAPICodePath("../common/api")
	//设置生成的文件后缀
	util.SetGenerateFileNameSuffix("rpc")
	//根据接口生成对应的rpc结构
	util.GeneratorRPC[services.IPropRPCService](prop.ModuleName, prop.Version, "propservice", "prop")
}
