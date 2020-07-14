package main

import (
	"cor-server-go/cor"
	"fmt"
	"time"
)

/**
 * Demo Only
 * 请使用自己的 AuthKey 和 SecretKey
 * AuthKey: WbykCN****8pwd3
 * SecretKey: 23423****dfsfs
 * ChannelName: 1234567****shao
 * UserId: 159****91766-10**26
 * ExpireTime: 1594704891
 * AuthToken: eyJ0b2tlbiI6ImZlNTdkNzUzMDI1ZDIwMzdlNGZjMDdhMmRkYTBhZmMxIiwidGltZXN0YW1wIjoxNTk0NzA0ODkxLCJ1c2VyYWNjb3VudCI6InJlc2VydmUiLCJyb2xlIjoicmVzZXJ2ZSJ9  lADPXbcGTTuoSBvn
 */
func main() {
	client := cor.Client{
		AuthKey:   "WbykCN****8pwd3",
		SecretKey: "23423****dfsfs",
	}
	expiredAt := time.Now().Add(60 * time.Minute).Unix()
	authToken, err := client.GetToken("1234567****shao", "159****91766-10**26", expiredAt)
	if err != nil {
		panic(err)
	}

	fmt.Println(expiredAt)
	fmt.Println(authToken)
}
