package tools

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken(12)
	t.Logf("token: %v", token)
	t.Logf("err: %v", err)
}

func TestCheckToke(t *testing.T) {
	token, _ := CreateToken(12)
	flag, id, err := CheckToke(token)
	t.Logf("flag: %v,id: %v,err:%v", flag, id, err)
}
