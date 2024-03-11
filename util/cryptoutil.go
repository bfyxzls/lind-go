package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5String(value string) string {
	data := []byte(value)                     // 待计算哈希的数据
	hash := md5.Sum(data)                     // 计算 MD5 哈希值
	hashString := hex.EncodeToString(hash[:]) // 将哈希值转换为字符串
	return hashString
}
