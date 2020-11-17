package common

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func GenetateAuth () string {
	h := md5.New()
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	h.Write([]byte(timeStr)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
