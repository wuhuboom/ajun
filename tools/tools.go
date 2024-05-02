package tools

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

var (
	SuccessBackCode = 200
	ErrorBackCode   = -100
)

func ReturnResponse(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK,
		gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
	return
}

// MD5hash  md5 加密
func MD5hash(password string) string {
	h := md5.Sum([]byte(password))  //用于计算哈希值并返回结果的操作
	return hex.EncodeToString(h[:]) //将字节数组转换为十六进制字符串的操作
}

// GenerateRandomString 生成随机字符串
func GenerateRandomString(length int) []byte {
	var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var strByteLen = len(strByte)
	bytes := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		bytes[i] = strByte[r.Intn(strByteLen)]
	}

	return bytes
}
