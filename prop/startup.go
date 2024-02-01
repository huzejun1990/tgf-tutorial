// @Author huzejun 2024/1/31 23:49:00
package prop

import (
	"github.com/thkhxm/tgf"
	"github.com/thkhxm/tgf/rpc"
)

func Startup() {
	r := rpc.NewRPCServer().
		WithRandomServicePort(8021, 8030).
		WithCache(tgf.CacheModuleClose).
		WithService(NewService()).
		WithGameConfig("../Common/conf/js").
		Run()
	<-r
}
