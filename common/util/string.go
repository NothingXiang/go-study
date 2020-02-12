/*
 * 工具包
 */
package util

import "strings"

// 判断是否空字符串，是返回true,否则返回false
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
