package helper

import (
	"crypto/md5"
	"encoding/hex"
	"unicode/utf8"
)

// InArray 判断元素是否在数组中
func InArray(needle interface{}, hystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hystack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hystack.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range hystack.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}

// Md5Encrypt md5加密
func Md5Encrypt(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

// MbStrLen 获取字符串长度
func MbStrLen(str string) int {
	return utf8.RuneCountInString(str)
}
