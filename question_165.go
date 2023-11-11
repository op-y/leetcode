package leetcode

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1s := strings.Split(version1, ".")
	v2s := strings.Split(version2, ".")
	l1 := len(v1s)
	l2 := len(v2s)
	for i := 0; i < l1 || i < l2; i++ {
		n1, n2 := 0, 0
		if i < l1 {
			n1, _ = strconv.Atoi(v1s[i])
		}
		if i < l2 {
			n2, _ = strconv.Atoi(v2s[i])
		}
		if n1 > n2 {
			return 1
		}
		if n1 < n2 {
			return -1
		}
	}
	return 0
}
