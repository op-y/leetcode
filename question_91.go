package leetcode

func numDecodings(s string) int {
	l := len(s)
	dp := make([]int, l+1)
	dp[0] = 1
	for i:=1; i<=l; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			dp[i] += dp[i-2]
		}
	}
   	return dp[l] 
}


