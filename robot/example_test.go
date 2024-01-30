// @Author huzejun 2024/1/30 21:44:00
package robot_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/huzejun1990/tgf/tgf-tutorial/common/pb"
	"github.com/thkhxm/tgf/robot"
	"testing"
)

func CreateRobot() robot.IRobot {
	rb := robot.NewRobotWs("tgf")
	rb.Connect("127.0.0.1:8443")
	return rb
}

func TestHelloWorld(t *testing.T) {
	rb := CreateRobot()
	req := &pb.HelloWorldRequest{Name: "tgf"}

	rb.RegisterCallbackMessage("hall.HelloWorld", func(i robot.IRobot, bytes []byte) {
		resp := &pb.HelloWorldResponse{}
		proto.Unmarshal(bytes, resp)
		t.Log(resp.Message)
	})

	rb.SendMessage("hall", "HelloWorld", req)

	select {}
}
