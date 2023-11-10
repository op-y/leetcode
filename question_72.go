package leetcode

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	if (l1 * l2) == 0 {
		return l1 + l2
	}

	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		r := make([]int, l2+1)
		dp[i] = r
	}
	for i := 0; i <= l1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			left_down := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				left_down += 1
			}
			dp[i][j] = min(left, min(down, left_down))
		}
	}
	return dp[l1][l2]
}
