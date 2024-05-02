package leetcode

func shortestPalindrome2(s string) string {
	n := len(s)
	nexts := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = nexts[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		nexts[i] = j
	}

	j := 0
	for i := n - 1; i >= 0; i-- {
		for j > 0 && s[i] != s[j] {
			j = nexts[j-1]
		}
		if s[i] == s[j] {
			j++
		}
	}

	add := ""
	if j != n {
		add = s[j:]
	}
	b := []byte(add)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b) + s
}
