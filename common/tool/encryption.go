package tool

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// EnCoder 密码加密
func EnCoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
