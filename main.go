package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/lwaly/tars_gateway_test_client/common"
)

type User struct {
	uid        uint32
	phone      string
	passwd     string
	privatekey []byte
}

var mapUser map[string]*User
var wg sync.WaitGroup

func main() {
	// testEncrypt()
	// Token_auth("token")
	strList := common.Conf.GetKeyList("user")
	mapUser = make(map[string]*User)
	//读取登录账号
	for _, value := range strList {
		if str, err := common.Conf.GetValue("user", value); nil == err {
			ss := strings.Split(str, ",")
			temp := User{phone: ss[0], passwd: ss[1]}
			mapUser[temp.phone] = &temp
		}
	}
	begin := time.Now().Unix()
	for _, user := range mapUser {
		wg.Add(1)
		go func(user *User) {
			defer wg.Done()
			//连接网关
			conn, err := net.Dial("tcp", "192.168.0.20:10008")
			if nil != err {
				fmt.Print("fail to connect", err)
				return
			}
			//登录
			user.privatekey, user.uid = UserLogin(conn, 1, user.phone, user.passwd)
			for i := 0; i < 10000; i++ {
				//循环获取手机区号
				ZoneGet(conn, user.phone)
				time.Sleep(time.Duration(3) * time.Second)
			}
		}(user)
	}
	wg.Wait()
	fmt.Print(time.Now().Unix() - begin)
}
