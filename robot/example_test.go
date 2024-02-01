// @Author huzejun 2024/1/30 21:44:00
package robot_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/huzejun1990/tgf/tgf-tutorial/common/api"
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

	rb.RegisterCallbackMessage(api.HelloWorld.MessageType, func(i robot.IRobot, bytes []byte) {
		resp := &pb.HelloWorldResponse{}
		proto.Unmarshal(bytes, resp)
		t.Log(resp.Message)
	}).RegisterCallbackMessage(api.Login.MessageType, func(i robot.IRobot, bytes []byte) {
		resp := &pb.LoginResponse{}
		proto.Unmarshal(bytes, resp)
		if resp.Success {
			i.Send(api.LoadUserData.MessageType, &pb.HelloWorldRequest{Name: "tgf"})
		} else {
			t.Log("login fail")
		}
	})

	//rb.SendMessage("hall", "HelloWorld", req)
	rb.Send(api.Login.MessageType, &pb.LoginRequest{
		Account:  "admin",
		Password: "123",
	})

	select {}
}

func TestRPCRobot(t *testing.T) {
	rb := CreateRobot()

	rb.RegisterCallbackMessage(api.Login.MessageType, func(i robot.IRobot, bytes []byte) {
		resp := &pb.LoginResponse{}
		proto.Unmarshal(bytes, resp)
		if resp.Success {
			i.Send(api.LoadUserData.MessageType, &pb.LoadUserDataRequest{})
		} else {
			t.Log("login fail")
		}
	}).RegisterCallbackMessage(api.LoadUserData.MessageType, func(i robot.IRobot, bytes []byte) {
		resp := &pb.LoadUserDataResponse{}
		proto.Unmarshal(bytes, resp)
		t.Log(resp.Name, resp.Name, resp.PropCount)
	})

	rb.Send(api.Login.MessageType, &pb.LoginRequest{
		Account:  "admin",
		Password: "123",
	})

	select {}
}
