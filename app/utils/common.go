package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// GetOffset 自动计算Offset
func GetOffset(page int, perPage int) int {
	if page < 1 {
		page = 1
	}
	return (page - 1) * perPage
}

// Md5Password 生成密码
func Md5Password(password string, salt string) string {
	srcCode := md5.Sum([]byte(password + salt))
	return fmt.Sprintf("%x", srcCode)
}

// GetRandomString 生成随机字符串,strType,0：全部，1：纯字母，2：纯数字
func GetRandomString(len int, strType int64) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var b []byte
	switch strType {
	case 0:
		b = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	case 1:
		b = []byte("0123456789")
	case 2:
		b = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	}

	result := make([]byte, len)
	for i := 0; i < len; i++ {
		result[i] = b[r.Int31()%62]
	}
	return string(result)
}

// GetIPAddress 获取IP地址
func GetIPAddress(req *http.Request) string {
	forwarded := req.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}
	ip := req.RemoteAddr
	// RemoteAddr 可能返回 IP 和端口号，例如 "192.168.1.1:12345"，这里我们简单地取 IP 部分
	idx := strings.LastIndex(ip, ":")
	if idx != -1 {
		ip = ip[:idx]
	}
	return ip
}

// VerifyTimeStr 验证时间格式
func VerifyTimeStr(layout string, value string) bool {
	_, err := time.ParseInLocation(layout, value, time.Local)
	if err != nil {
		return false
	}
	return true
}
