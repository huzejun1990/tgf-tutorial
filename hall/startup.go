// @Author huzejun 2024/1/30 17:16:00
package hall

import (
	"github.com/thkhxm/tgf/rpc"
)

func Startup() {
	r := rpc.NewRPCServer().
		WithService(NewService()).
		WithGatewayWS("8443", "tgf", nil).
		//WithCache(tgf.CacheModuleClose).
		WithRandomServicePort(8010, 8020).
		WithWhiteService("Login").
		Run()

	<-r
}
