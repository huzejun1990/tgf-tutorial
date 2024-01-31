// @Author huzejun 2024/1/31 21:25:00
package services

import (
	"context"
	"github.com/huzejun1990/tgf/tgf-tutorial/common/model"
)

type IPropRPCService interface {
	GetUserPropCount(ctx context.Context, args *model.GetUserPropArgs, reply *model.GetUserPropReply) (err error)
}
