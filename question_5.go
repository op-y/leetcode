package leetcode

func longestPalindrome(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }

    ans := ""
    maxLen := 1
    dp := make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, n)
    }

    // 初始化对角线
    for i:=0; i < n; i++ {
        dp[i][i] = true
    }

    // 按列计算
    for j := 1; j < n; j++ {
        for i := 0; i < j; i++ {
            if s[i] != s[j] {
                dp[i][j] = false
            } else {
                if j - i < 3 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }
            }

            if dp[i][j] && j - i + 1 > maxLen {
                maxLen = j - i + 1
                ans = s[i:i+maxLen]
            }
        }
    }

    return ans
}
