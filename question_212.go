package leetcode

import (
	"math"
)

// 这里只使用了回溯 官方题解用了字典树
// 已经出现的词语可以从wordMap删除实现另一种去重
func findWords212(board [][]byte, words []string) []string {
	lmax, lmin := 0, math.MaxInt
	wordMap := make(map[string]bool)
	for _, word := range words {
		wordMap[word] = true
		if len(word) < lmin {
			lmin = len(word)
		}
		if len(word) > lmax {
			lmax = len(word)
		}
	}

	ans := map[string]bool{}
	m, n := len(board), len(board[0])
	flags := make([][]bool, m)
	for i := 0; i < m; i++ {
		flags[i] = make([]bool, n)
	}

	lword := 0
	word := []byte{}
	var backoff func(int, int)
	backoff = func(x, y int) {
		flags[x][y] = true
		lword++
		word = append(word, board[x][y])

		if lword > lmax {
			flags[x][y] = false
			word = word[:len(word)-1]
			lword--
			return
		}

		if lword >= lmin && lword <= lmax {
			if _, ok := wordMap[string(word)]; ok {
				ans[string(word)] = true
			}
		}

		if x-1 >= 0 && !flags[x-1][y] {
			backoff(x-1, y)
		}
		if x+1 < m && !flags[x+1][y] {
			backoff(x+1, y)
		}
		if y-1 >= 0 && !flags[x][y-1] {
			backoff(x, y-1)
		}
		if y+1 < n && !flags[x][y+1] {
			backoff(x, y+1)
		}

		flags[x][y] = false
		word = word[:len(word)-1]
		lword--
		return

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backoff(i, j)
		}
	}

	result := []string{}
	for word, _ := range ans {
		result = append(result, word)
	}
	return result
}

//func main() {
//	board := [][]byte{
//		{'o', 'a', 'a', 'n'},
//		{'e', 't', 'a', 'e'},
//		{'i', 'h', 'k', 'r'},
//		{'i', 'f', 'l', 'v'},
//	}
//	words := []string{"oath", "pea", "eat", "rain"}
//	//board := [][]byte{
//	//	{'a', 'b'},
//	//	{'c', 'd'},
//	//}
//	//words := []string{"abcb"}
//	ans := findWords(board, words)
//	fmt.Printf("%v\n", ans)
//}
