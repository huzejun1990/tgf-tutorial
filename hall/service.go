// @Author huzejun 2024/1/30 17:17:00
package hall

import (
	"context"
	"fmt"
	"github.com/huzejun1990/tgf/tgf-tutorial/common/pb"
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
)

type service struct {
}

func (s *service) HelloWorld(ctx context.Context, args *rpc.Args[*pb.HelloWorldRequest], reply *rpc.Reply[*pb.HelloWorldResponse]) (err error) {
	//args.GetData() ->  *pb.HelloWorldRequest

	msg := fmt.Sprintf("hello world %s", args.GetData().GetName())
	reply.SetData(&pb.HelloWorldResponse{Message: msg})
	return
}

func (s *service) Login(ctx context.Context, args *rpc.Args[*pb.LoginRequest], reply *rpc.Reply[*pb.LoginResponse]) (err error) {
	log.DebugTag("hall", "hall login")
	return
}

func (s *service) GetName() string {
	return "hall"
}

func (s *service) GetVersion() string {
	return "1.0.0"
}

func (s *service) Startup() (bool, error) {
	log.DebugTag("hall", "hall startup")
	return true, nil
}

func (s *service) Destroy(sub rpc.IService) {
	log.DebugTag("hall", "hall destroy")
}

func NewService() rpc.IService {
	return &service{}
}
