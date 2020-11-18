package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
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

	mapUser = make(map[string]*User)
	// go tcpTest()
	httpTest()
	wg.Wait()
}

func tcpTest() {
	strList := common.Conf.GetKeyList("user")
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
			conn, err := net.Dial("tcp", "127.0.0.1:8480")
			if nil != err {
				fmt.Print("fail to connect", err)
				return
			}
			//登录
			user.privatekey, user.uid = UserLogin(conn, 1, user.phone, user.passwd)
			for i := 0; i < 10; i++ {
				//循环获取手机区号
				ZoneGet(conn, user.phone)
				time.Sleep(time.Duration(3) * time.Second)
			}
		}(user)
	}
	fmt.Print(time.Now().Unix() - begin)
}

func httpTest() {
	data := map[string]interface{}{
		"page_size":  3,
		"page_index": 1,
		"id":         1,
	}
	b, _ := json.Marshal(data)

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://127.0.0.1:18480/a/b/c", bytes.NewReader(b))
	if err != nil {
		fmt.Print(err)
		return
	}

	req.Header.Set("Token", Create_token(1, "ly"))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(string(body))
}
