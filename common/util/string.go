/*
 * 工具类
 */
package util

import (
	"strings"
)

func IsEmpty(s string) bool {
	if strings.TrimSpace(s); len(s) != 0 {
		return false
	}
	return true
}
