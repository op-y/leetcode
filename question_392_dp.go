package leetcode

func isSubsequenceDP(s string, t string) bool {
	n, m := len(s), len(t)

	dp := make([][26]int, m+1)
	for i := 0; i < 26; i++ {
		dp[m][i] = m // 最后一个字母往后都记为m 无法到达
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j+'a') {
				dp[i][j] = i
			} else {
				dp[i][j] = dp[i+1][j] // 由于是从后往前处理这里存的是尽可能远的位置非命中位置
			}
		}
	}
	add := 0
	for i := 0; i < n; i++ {
		if dp[add][int(s[i]-'a')] == m {
			return false // m为不可达
		}
		add = dp[add][int(s[i]-'a')] + 1
	}
	return true
}
