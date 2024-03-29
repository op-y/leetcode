package leetcode

// 官方题解 KMP算法
func shortestPalindrome(s string) string {
	n := len(s)

	fail := make([]int, n)
	for i := 0; i < n; i++ {
		fail[i] = -1
	}

	// DFA计算
	for i := 1; i < n; i++ {
		j := fail[i-1]
		for j != -1 && s[j+1] != s[i] {
			j = fail[j]
		}
		if s[j+1] == s[i] {
			fail[i] = j + 1
		}
	}

	// 获取最长的保留长度
	best := -1
	for i := n - 1; i >= 0; i-- {
		for best != -1 && s[best+1] != s[i] {
			best = fail[best]
		}
		if s[best+1] == s[i] {
			best++
		}
	}

	// 反转不是回文前缀的部分
	add := ""
	if best != n-1 {
		add = s[best+1:]
	}
	b := []byte(add)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}

	return string(b) + s
}

//func main() {
//	s := "aacecaaa"
//	r := shortestPalindrome(s)
//	fmt.Println(r)
//}
