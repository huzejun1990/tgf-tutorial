// @Author huzejun 2024/1/31 23:28:00
package model

type GetUserPropArgs struct {
	PropId string
}

type GetUserPropReply struct {
	Name  string
	Count int32
}
