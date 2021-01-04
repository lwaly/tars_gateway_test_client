package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	proto "github.com/golang/protobuf/proto"
	"github.com/lwaly/tars_gateway_test_client/user_proto"
)

var secret = "test"
var privatekey = []byte("")

type Claims struct {
	Uid  uint64 `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

const (
	OK                 = 0
	ERR_UNKNOWN        = 10000
	ERR_INPUT_DATA     = 10001
	ERR_DATABASE       = 10002
	ERR_DUP_USER       = 10003
	ERR_NO_USER        = 10004
	ERR_PASS           = 10005
	ERR_NO_USE_RPASS   = 10006
	ERR_NO_USER_CHANGE = 10007
	ERR_INVALID_USER   = 10008
	ERR_OPEN_FILE      = 10009
	ERR_WRIT_EFILE     = 10010
	ERR_SYSTEM         = 10011
	ERR_EXPIRED        = 10012
	ERR_PERMISSION     = 10013
)

const (
	ErrInputData    = "数据输入错误"
	ErrDatabase     = "数据库操作错误"
	ErrDupUser      = "用户信息已存在"
	ErrNoUser       = "用户信息不存在"
	ErrPass         = "密码不正确"
	ErrNoUserPass   = "用户信息不存在或密码不正确"
	ErrNoUserChange = "用户信息不存在或数据未改变"
	ErrInvalidUser  = "用户信息不正确"
	ErrOpenFile     = "打开文件出错"
	ErrWriteFile    = "写文件出错"
	ErrSystem       = "操作系统错误"
	ErrUnknown      = "未知错误"
)

func init() {
}

func To_md5(userid uint64, expireToken int64) (decode []byte) {
	md5Ctx := md5.New()
	encode := fmt.Sprintf("%s%d%d", secret, userid, expireToken)
	md5Ctx.Write([]byte(encode))
	cipherStr := md5Ctx.Sum([]byte(secret))
	return []byte(base64.StdEncoding.EncodeToString(cipherStr))
}

func Create_token(id uint64, name string) string {
	expireToken := time.Now().Add(time.Hour * 24 * 30).Unix()
	claims := Claims{
		//info.Appid,
		id,
		name,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	}

	// Create the token using your claims
	c_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := c_token.SignedString(To_md5(id, expireToken))

	return signedToken
}

func Token_auth(signedToken string) (uint64, string, error) {
	s := strings.Split(signedToken, ".")
	if 3 != len(s) {
		return 0, "", errors.New(ErrDupUser)
	}

	claims := Claims{}

	if b, err := jwt.DecodeSegment(s[1]); nil != err {
		return 0, "", err
	} else {
		if err := json.Unmarshal(b, &claims); nil != err {
			return 0, "", err
		}
	}
	fmt.Printf("%v %s", claims, secret)
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return To_md5(claims.Uid, claims.ExpiresAt), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("%d %d", claims.Uid, claims.ExpiresAt)
		return claims.Uid, claims.Name, err
	}
	return 0, "", err
}

func SendMsg(conn net.Conn, phone string, b []byte, server, cmd, cacheis int) (ret, pri []byte) {
	var err error
	v, ok := mapUser[phone]
	//判断是否加载秘钥，进行消息加密
	encrypt := uint32(2)
	if ok && 0 != len(v.privatekey) {
		publickey, _ := LoadPublicKeyBytes(v.privatekey)
		b, err = EncryptPkcs(b, publickey)
	} else {
		encrypt = 3
	}
	//业务体封装
	body := user_proto.MsgBody{Body: b}
	b, _ = proto.Marshal(&body)
	head := user_proto.MsgHead{Version: 1, App: 1, Server: uint32(server), Servant: uint32(cmd),
		Seq: 1, RouteId: 1, BodyLen: uint32(len(b)), Encrypt: encrypt, CacheIs: uint32(cacheis)}
	outHeadRsp, err := proto.Marshal(&head)
	if err != nil {
		fmt.Printf("faile to Marshal msg.err: %v", err)
		return
	}
	outHeadRsp = append(outHeadRsp, b...)
	//发送
	n, err := conn.Write(outHeadRsp)
	if err != nil {
		fmt.Printf("Write msg.err: %v", err)
		return
	}
	var data1 [49]byte
	//读取头
	n, err = conn.Read(data1[:])
	if err != nil {
		fmt.Printf("Read msg.err:%v ", err)
		return
	}

	var head1 user_proto.MsgHead
	err = proto.Unmarshal(data1[:], &head1)
	if err != nil {
		fmt.Printf("faile to Unmarshal msg.err: %v", err)
		return
	}

	buf := make([]byte, head1.GetBodyLen())
	total := 0
AGAIN:
	//读取体
	n, err = conn.Read(buf[total:])
	if (0 == n) || (nil != err) {
		fmt.Printf("faile to Unmarshal msg.err:%d %v ", n, err)
	}
	total += n
	if uint32(total) < head1.GetBodyLen() {
		goto AGAIN
	}
	var msgBody user_proto.MsgBody
	err = proto.Unmarshal(buf[:], &msgBody)
	if err != nil {
		fmt.Printf("faile to Unmarshal msg.err: %v", err)
		return
	}
	//判断是否是验证消息，解析秘钥
	if 0 != len(msgBody.GetExtend()) {
		privatekey = msgBody.GetExtend()
		pri = msgBody.GetExtend()
	}

	return msgBody.GetBody(), pri
}
