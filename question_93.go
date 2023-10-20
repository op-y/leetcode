package leetcode

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	result := []string{}

	count := 0 // 记录获取了几段
	seg := []string{}

	var getIP func(string, int)
	getIP = func(s string, pos int) {
		if len(s)-pos > (4-count)*3 {
			return
		}
		if len(s)-pos < (4-count)*1 {
			return
		}

		if count == 4 && pos == len(s) {
			ip := strings.Join(seg, ".")
			result = append(result, ip)
			return
		}

		n1, _ := strconv.Atoi(s[pos : pos+1])
		if pos < len(s) {
			count++
			seg = append(seg, s[pos:pos+1])
			getIP(s, pos+1)
			seg = seg[:len(seg)-1]
			count--
		}
		if n1 != 0 && pos+1 < len(s) {
			count++
			seg = append(seg, s[pos:pos+2])
			getIP(s, pos+2)
			seg = seg[:len(seg)-1]
			count--
		}
		if n1 != 0 && pos+2 < len(s) {
			n3, _ := strconv.Atoi(s[pos : pos+3])
			if n3 > 255 {
				return
			}
			count++
			seg = append(seg, s[pos:pos+3])
			getIP(s, pos+3)
			seg = seg[:len(seg)-1]
			count--
		}
	}
	getIP(s, 0)
	return result
}
