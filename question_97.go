package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	ls1, ls2, ls3 := len(s1), len(s2), len(s3)
	if ls1 + ls2 != ls3 {
		return false
	}
	dp := make([][]bool, ls1+1)
	for i:=0; i<=ls1; i++ {
		dp[i] = make([]bool, ls2+1)
	}
	dp[0][0] = true
	for i:=0; i<=ls1; i++ {
		for j:=0; j<=ls2; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && (s1[i-1] == s3[p])	
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && (s2[j-1] == s3[p])	
			}
		}
	}
	return dp[ls1][ls2]
}
