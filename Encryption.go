package tools

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5Encryption MD5加盐加密
func MD5Encryption(password, salt string) string {
	m := md5.New()
	m.Write([]byte(password + salt))
	return hex.EncodeToString(m.Sum(nil))
}
