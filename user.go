package main

import (
	"fmt"
	"net"

	"github.com/lwaly/tars_gateway_test_client/user_proto"

	proto "github.com/golang/protobuf/proto"
)

func UserLogin(conn net.Conn, t uint32, phone string, passwd string) ([]byte, uint32) {
	//发送消息组装
	req := user_proto.LoginReq{Phone: []byte(phone), VerifyInfo: []byte(passwd), Type: 1}
	b, _ := proto.Marshal(&req)
	//消息发送
	b, pri := SendMsg(conn, phone, b, 2, int(user_proto.ECmd_E_LOGIN_REQ))
	//解析返回
	rsp := user_proto.LoginRsp{}
	proto.Unmarshal(b, &rsp)
	return pri, rsp.GetId()
}

func ZoneGet(conn net.Conn, phone string) {
	req := user_proto.ZoneGetReq{}
	b, _ := proto.Marshal(&req)
	b, _ = SendMsg(conn, phone, b, 2, int(user_proto.ECmd_E_ZONE_GET_REQ))
	rsp := user_proto.ZoneGetRsp{}
	proto.Unmarshal(b, &rsp)
	fmt.Printf("%v %s\n", rsp, string(rsp.GetList()[0].Country))
}
