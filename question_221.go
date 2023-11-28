package leetcode

// 动态规划 计算最长的边
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	lmax := 0
	for i:=0; i<len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j:=0; j<len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				lmax = 1
			}
		}
	}

	for i := 1; i<len(matrix); i++ {
		for j := 1; j<len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > lmax {
					lmax = dp[i][j]
				}
			}
		}
	}

	return lmax * lmax
}
