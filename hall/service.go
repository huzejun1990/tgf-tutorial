// @Author huzejun 2024/1/30 17:17:00
package hall

import (
	"context"
	"fmt"
	propservice "github.com/huzejun1990/tgf/tgf-tutorial/common/api/prop"
	"github.com/huzejun1990/tgf/tgf-tutorial/common/model"
	"github.com/huzejun1990/tgf/tgf-tutorial/common/pb"
	"github.com/huzejun1990/tgf/tgf-tutorial/common/services"
	"github.com/huzejun1990/tgf/tgf-tutorial/hall/internal"
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
	"github.com/thkhxm/tgf/util"
)

var (
	ModuleName                       = "hall"
	Version                          = "v1.0.0"
	_          services.IHallService = new(service)
)

type service struct {
	m *internal.Manager
}

func (s *service) UpdatePassword(ctx context.Context, args *rpc.Args[*pb.LoginRequest], reply *rpc.Reply[*pb.LoginResponse]) (err error) {
	req := args.GetData()
	s.m.UpdatePassword(req.Account, req.Password)
	return
}

func (s *service) LoadUserData(ctx context.Context, args *rpc.Args[*pb.LoadUserDataRequest], reply *rpc.Reply[*pb.LoadUserDataResponse]) (err error) {
	// 用户登陆成功之后，请求大厅的LoadUserData接口，大厅节点通过rpc,获取Prop节点的道具信息
	//propId, _ := util.AnyToStr(rand.Int31n(10))
	propId, _ := util.AnyToStr(100033)
	rpcReply := &model.GetUserPropReply{}
	rpc.SendRPCMessage(ctx, propservice.GetUserPropCount.New(&model.GetUserPropArgs{PropId: propId}, rpcReply))
	log.DebugTag("hall", "load user data propId=%v propName=%s count=%v", propId, rpcReply.Name, rpcReply.Count)
	reply.SetData(&pb.LoadUserDataResponse{Name: rpcReply.Name, PropCount: rpcReply.Count})
	return
}

func (s *service) Login(ctx context.Context, args *rpc.Args[*pb.LoginRequest], reply *rpc.Reply[*pb.LoginResponse]) (err error) {
	var userId string
	var pbData = &pb.LoginResponse{Success: false}

	account, code := s.m.GetAccount(args.GetData().GetAccount(), args.GetData().GetPassword())
	if code == 0 {
		if account == nil {
			account = s.m.CreateAccount(args.GetData().GetAccount(), args.GetData().GetPassword(), util.GenerateSnowflakeId())
		}
		pbData.Success = true
		rpc.UserLogin(ctx, userId)
	} else {
		reply.SetCode(code)
	}
	pbData.UserId = account.UserId
	/*	if args.GetData().GetAccount() == "admin" && args.GetData().GetPassword() == "123" {
			userId = util.GenerateSnowflakeId()
			rpc.UserLogin(ctx, userId)
			pbData = &pb.LoginResponse{Success: true}
		} else {
			pbData = &pb.LoginResponse{Success: false}

		}*/
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
	s.m = internal.NewManager()
	return true, nil
}

func (s *service) Destroy(sub rpc.IService) {
	log.DebugTag("hall", "hall destroy")
}

func NewService() rpc.IService {
	return &service{}
}
