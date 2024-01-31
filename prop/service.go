// @Author huzejun 2024/1/31 23:53:00
package prop

import (
	"context"
	"github.com/huzejun1990/tgf/tgf-tutorial/common/model"
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
	"github.com/thkhxm/tgf/util"
	"math/rand"
)

var (
	ModuleName = "prop"
	Version    = "v1.0.0"
)

type service struct {
	rpc.Module
	propCountCache map[string]int32
}

func (s *service) GetUserPropCount(ctx context.Context, args *model.GetUserPropArgs, reply *model.GetUserPropReply) (err error) {
	userId := rpc.GetUserId(ctx)
	reply.Count = s.propCountCache[args.PropId]
	log.DebugTag("prop", "get %s user %s prop count ", userId, args.PropId, reply.Count)

	return
}

func (s *service) GetName() string {
	return ModuleName
}

func (s *service) GetVersion() string {
	return Version
}

func (s *service) Startup() (bool, error) {
	s.propCountCache = make(map[string]int32)
	r := rand.New(rand.NewSource(12345))
	for i := 0; i < 10; i++ {
		propId, _ := util.AnyToStr(i)
		s.propCountCache[propId] = r.Int31n(100)
	}
	return true, nil
}

func (s *service) Destroy(sub rpc.IService) {
}

func NewService() rpc.IService {
	return &service{}
}
