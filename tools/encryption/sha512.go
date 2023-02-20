package encryption

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashEncode(str string) string {
	SHA := sha512.Sum512([]byte(str))
	res := hex.EncodeToString(SHA[:])
	return res
}
