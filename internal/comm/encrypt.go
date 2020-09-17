package comm

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var randSource = rand.NewSource(time.Now().UnixNano())

// 加密
func EncryptMD5(pwdStr string) string {
	ranS := RandomStr(len(pwdStr))
	needEncrypt := []byte(pwdStr + ranS)
	encryptStr := md5.Sum(needEncrypt)
	encryptStr = md5.Sum([]byte(fmt.Sprintf("%x", encryptStr)))
	return fmt.Sprintf("%x", encryptStr)
}

// 生成随机字符串
func RandomStr(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, randSource.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSource.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
