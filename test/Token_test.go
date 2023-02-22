package main

import (
	"awesomeProject/dou-yin/common/token"
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, err := tools.CreateToken(1)
	t.Logf("token: %v", token)
	t.Logf("err: %v", err)
}

func TestCheckToke(t *testing.T) {
	token, _ := tools.CreateToken(1)
	flag, id, err := tools.CheckToke(token)
	t.Logf("flag: %v,id: %v,err:%v", flag, id, err)
}
