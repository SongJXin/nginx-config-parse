package util

import "fmt"

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

// compareAndColorize 比较两个字符串并返回带颜色的字符串
func CompareAndColorize(str1, str2 string) string {
	if str1 == str2 {
		return fmt.Sprintf("%s%s%s", Green, str1, Reset)
	}
	return fmt.Sprintf("%s%s%s", Red, str2, Reset)
}
