package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		}
		if i > 0 && dp[i-1][0] == 0 {
			dp[i][0] = 0
		}
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = 1
		}
		if j > 0 && dp[0][j-1] == 0 {
			dp[0][j] = 0
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
