package leetcode

//func ladderLength(beginWord string, endWord string, wordList []string) int {
//	// 标记已经使用的单词
//	flags := make([]bool, len(wordList))
//	count, min := 1, 0
//	var backoff func(string)
//	backoff = func(word string) {
//		// 完成转换
//		if word == endWord {
//			if min == 0 || count < min {
//				min = count
//			}
//			return
//		}
//
//		for idx, w := range wordList {
//			if flags[idx] {
//				// 该位置已经被使用过则跳过
//				continue
//			}
//			if canTrans(word, w) {
//				count++
//				flags[idx] = true
//				backoff(w)
//				// 回溯
//				flags[idx] = false
//				count--
//			}
//		}
//		// 全部单词都使用完
//		return
//	}
//	backoff(beginWord)
//	return min
//}
//
//func canTrans(from, to string) bool {
//	l := len(from)
//	found := false
//	for i := 0; i < l; i++ {
//		if from[i] == to[i] {
//			continue
//		}
//		if !found {
//			found = true
//		} else {
//			// 超过一个字符的差异
//			return false
//		}
//	}
//	return found
//}

//func main() {
//	beginWord := "hit"
//	endWord := "cog"
//	//wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
//	wordList := []string{"hot", "dot", "dog", "lot", "log"}
//	min := ladderLength(beginWord, endWord, wordList)
//	fmt.Printf("%d\n", min)
//}
