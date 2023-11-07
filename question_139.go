package leetcode

//func wordBreak(s string, wordDict []string) bool {
//	var backoff func(start int) bool
//	backoff = func(start int) bool {
//		if start == len(s) {
//			return true
//		}
//		for _, word := range wordDict {
//			l := len(word)	
//			if start + l > len(s) {
//				continue
//			}
//			if word == s[start:start+l] {
//				if ok := backoff(start+l); ok {
//					return ok
//				}
//			}
//		}
//		return false	
//	}
//	return backoff(0)	
//}

func wordBreak(s string, wordDict []string) bool {
    wordDictSet := make(map[string]bool)
    for _, w := range wordDict {
        wordDictSet[w] = true
    }
    dp := make([]bool, len(s) + 1)
    dp[0] = true
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordDictSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[len(s)]
}
