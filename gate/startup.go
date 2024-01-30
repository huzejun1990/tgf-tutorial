// @Author huzejun 2024/1/29 22:56:00
package gate

import (
	"github.com/thkhxm/tgf"
	"github.com/thkhxm/tgf/rpc"
)

func Startup() {
	r := rpc.NewRPCServer().
		WithGatewayWS("8443", "/tgf", nil).
		WithCache(tgf.CacheModuleClose).
		Run()

	<-r
}
