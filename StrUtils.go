package utils

import "strings"

//检测是否是空串
func CheckStrIsEmpty(str string) bool {
	str = strings.TrimSpace(str)
	return len(str) < 1
}
