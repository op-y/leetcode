package leetcode

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	list := strings.Split(s, " ")
	rlist := []string{}
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] != "" {
			rlist = append(rlist, list[i])
		}
	}
	return strings.Join(rlist, " ")
}
