package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// jwt全称 json web token
// 是一种后端不用存储的前端身份验证工具 说白了就是不用再存储到数据库来验证身份
// token 理解成标记
// 分为三个部分 header头部 claims内容 signature签名/加密串
// jwt的显著特征就是认票不认人，认串不认人
// go get github.com/dgrijalva/jwt-go/cmd/jwt  要先下载这个包

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {

	mySigningKey := []byte("fangguowoba")

	//包含结构去创建一个jwt
	//SigningMethod是加密方法
	//claims 是实现claims一个结构体   常用的是MapClaims实现，就是map  还有最常用的StandardClaims实现,就是一个结构体
	//传入的参数两个 一个是加密方式，一般就使用RS256(非对称加密)，为了Demo方便，使用HS256 一个是结构体，要实例化
	c := MyClaims{
		Username: "gzy",
		StandardClaims: jwt.StandardClaims{
			//jwt.StandardClaims是一个结构体
			//type StandardClaims struct {
			//	Audience  string `json:"aud,omitempty"`
			//	ExpiresAt int64  `json:"exp,omitempty"`  过期时间
			//	Id        string `json:"jti,omitempty"`
			//	IssuedAt  int64  `json:"iat,omitempty"`
			//	Issuer    string `json:"iss,omitempty"`  签发人
			//	NotBefore int64  `json:"nbf,omitempty"`  开始生效时间
			//	Subject   string `json:"sub,omitempty"`
			//}

			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 5,
			Issuer:    "gzy",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, e := t.SignedString(mySigningKey)
	if e != nil {
		fmt.Printf("%s", e)
	}
	fmt.Println(s)
	//下面的就是生成的秘钥 每次都不是一样的
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2Vybm
	//FtZSI6Imd6eSIsImV4cCI6MTY5NjM5NzQ1MiwiaXNzIjoiZ3p5
	//IiwibmJmIjoxNjk2Mzk3Mzg3fQ.G03jkT4TH9GGYtPIFDAk_bKJsFkNQedJomDZ5Q5Rf9A

	time.Sleep(6 * time.Second)

	token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Printf("%s", err) //token is expired by 1s

		fmt.Println()
	}
	fmt.Println(token)
	//&{eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXV
	// CJ9.eyJ1c2VybmFtZSI6Imd6eSIsImV4cCI6MTY5NjM5NzU
	//0MCwiaXNzIjoiZ3p5IiwibmJmIjoxNjk2Mzk3NDc1fQ.eA
	//RvYr6CXS_U3u3gP6OL9yix0n4V-5w5PkYLn60OhEk
	//0xc000008090 map[alg:HS256 typ:JWT] 0xc000086000 eARvYr6CXS_U3u3gP6OL9yix0n4V-5w5PkYLn60OhEk false}

	fmt.Println(token.Claims.(*MyClaims).Username)
}
