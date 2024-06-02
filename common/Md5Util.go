package common

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 大写
func Md5EncodeUpper(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// salt加密
func Md5EncodeWithSalt(password string, salt string) string {
	return Md5Encode(password + salt)
}

// 解密
func Md5ValidateUpper(password string, encoded string) bool {
	return Md5EncodeUpper(password) == encoded
}

// salt解密
func Md5ValidateWithSalt(password string, salt string, encoded string) bool {
	return Md5Encode(password+salt) == encoded
}
