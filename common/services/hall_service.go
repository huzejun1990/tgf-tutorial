// @Author huzejun 2024/1/30 20:08:00
package services

import (
	"context"
	"github.com/huzejun1990/tgf/tgf-tutorial/common/pb"
	"github.com/thkhxm/tgf/rpc"
)

type IHallService interface {
	Login(ctx context.Context, args *rpc.Args[*pb.LoginRequest], reply *rpc.Reply[*pb.LoginResponse]) (err error)
	HelloWorld(ctx context.Context, args *rpc.Args[*pb.HelloWorldRequest], reply *rpc.Reply[*pb.HelloWorldResponse]) (err error)
}
