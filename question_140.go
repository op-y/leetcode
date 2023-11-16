package leetcode

import (
	"strings"
)

func wordBreak140(s string, wordDict []string) []string {
	wordMap := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = true
	}

	l := len(s)
	tbl := make([][]bool, l+1)
	for i := 0; i < len(tbl); i++ {
		tbl[i] = make([]bool, l+1)
	}
	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			if wordMap[s[j:i]] {
				tbl[j][i] = true
			}
		}
	}

	ans := []string{}
	dl := []string{}
	var check func(int)
	check = func(start int) {
		if start == l {
			ans = append(ans, strings.Join(dl, " "))
		}
		for end := start + 1; end <= l; end++ {
			if tbl[start][end] {
				dl = append(dl, s[start:end])
				check(end)
				dl = dl[:len(dl)-1]
			}
		}
	}
	check(0)

	return ans
}

//func main() {
//	//s := "catsanddog"
//	//wordDict := []string{"cat", "cats", "and", "sand", "dog"}
//	//s := "pineapplepenapple"
//	//wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
//	s := "catsandog"
//	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
//	wordBreak(s, wordDict)
//	return
//}
