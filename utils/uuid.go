package utils

import (
	"errors"
	"math/rand"
	"time"
)

func GenRandomXUIDByClientID(xuidLen int) (string,error)  {
	res := ""
	baseStr := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< xuidLen; i++ {
		position := rand.Intn(len(baseStr))
		char := baseStr[position]
		res += string(char)
	}
	if len(res) == xuidLen {
		return res,nil
	}
	return res,errors.New("获取xuid失败")
}
