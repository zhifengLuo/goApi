package library

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	crypt := md5.New()
	crypt.Write([]byte(str))
	return hex.EncodeToString(crypt.Sum(nil))
}
