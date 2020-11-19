package common

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego"
)

// GenetateAuth ...
func GenetateAuth() string {
	h := md5.New()
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	h.Write([]byte(timeStr)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return "Bearer" + " " + hex.EncodeToString(cipherStr)
}

// DetectGRPC ...
func DetectGRPC() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content) + beego.AppConfig.String("httpport")
}
