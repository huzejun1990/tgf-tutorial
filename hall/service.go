// @Author huzejun 2024/1/30 17:17:00
package hall

import (
	"context"
	"fmt"
	"github.com/huzejun1990/tgf/tgf-tutorial/common/pb"
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
	"github.com/thkhxm/tgf/util"
)

var (
	ModuleName = "hall"
	Version    = "v1.0.0"
)

type service struct {
}

func (s *service) Login(ctx context.Context, args *rpc.Args[*pb.LoginRequest], reply *rpc.Reply[*pb.LoginResponse]) (err error) {
	var userId string
	var pbData *pb.LoginResponse

	if args.GetData().GetAccount() == "admin" && args.GetData().GetPassword() == "123" {
		userId = util.GenerateSnowflakeId()
		rpc.UserLogin(ctx, userId)
		pbData = &pb.LoginResponse{Success: true}
	} else {
		pbData = &pb.LoginResponse{Success: false}

	}
	reply.SetData(pbData)
	return
}

func (s *service) HelloWorld(ctx context.Context, args *rpc.Args[*pb.HelloWorldRequest], reply *rpc.Reply[*pb.HelloWorldResponse]) (err error) {
	//args.GetData() ->  *pb.HelloWorldRequest

	msg := fmt.Sprintf("hello world %s", args.GetData().GetName())
	reply.SetData(&pb.HelloWorldResponse{Message: msg})
	return
}

func (s *service) GetName() string {
	return ModuleName
}

func (s *service) GetVersion() string {
	return Version
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
