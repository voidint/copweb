package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 计算数据的MD5摘要并返回其字符串形式的MD5值
func Md5String(data []byte) string {
	array := md5.Sum(data)
	return hex.EncodeToString(array[:])
	//return hex.EncodeToString(Md5Bytes(data))
}

// 计算数据的MD5摘要并返回其字节切片形式的MD5值
func Md5Bytes(data []byte) []byte {
	hash := md5.New()
	hash.Write(data)
	return hash.Sum(nil)
}
