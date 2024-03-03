package leetcode

import (
	"strconv"
)

func decodeString(s string) string {
	q := []byte{}

	for i := 0; i < len(s); i++ {
		// 遇到]出栈最近的一个重复段
		if isSegRight(s[i]) {
			seg := []byte{}
			num := []byte{}

			// 取最近一个括号对中的字符
			for !isSegLeft(q[len(q)-1]) {
				seg = append(seg, q[len(q)-1])
				q = q[0 : len(q)-1]
			}

			// 删除[字符
			q = q[0 : len(q)-1]

			// 取[之前的数字
			for len(q) > 0 && isDigit394(q[len(q)-1]) {
				num = append(num, q[len(q)-1])
				q = q[0 : len(q)-1]
			}
			for i, j := 0, len(num)-1; i < j; {
				num[i], num[j] = num[j], num[i]
				i++
				j--
			}
			times, _ := strconv.Atoi(string(num))

			// 重复后入栈
			for i := 0; i < times; i++ {
				for j := len(seg) - 1; j >= 0; j-- {
					q = append(q, seg[j])
				}
			}
		} else {
			q = append(q, s[i])
		}
	}

	return string(q)
}

func isDigit394(c byte) bool {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func isSegLeft(c byte) bool {
	if c == '[' {
		return true
	}
	return false
}

func isSegRight(c byte) bool {
	if c == ']' {
		return true
	}
	return false
}
