package leetcode

//import (
//	"fmt"
//)
//
//func ladderLength(beginWord string, endWord string, wordList []string) int {
//	// 结束词不在图中
//	if !inPath(endWord, wordList) {
//		return 0
//	}
//	// 将输入节点一起放入构造图
//	if !inPath(beginWord, wordList) {
//		wordList = append(wordList, beginWord)
//	}
//	graph := makeGraph(wordList)
//
//	path := shortestPath(beginWord, endWord, graph)
//	return path
//}
//
//func shortestPath(start, end string, graph map[string][]string) int {
//	var bfs func(string) int
//	bfs = func(node string) int {
//		count := 0
//		nodes := []string{}
//
//		cur := []string{node}
//		for len(cur) > 0 {
//			count++
//			next := []string{}
//			for _, cn := range cur {
//				nodes = append(nodes, cn)
//				if inPath(end, graph[cn]) {
//					return count + 1
//				}
//			}
//			for _, cn := range cur {
//				for _, nn := range graph[cn] {
//					if !inPath(nn, nodes) {
//						next = append(next, nn)
//					}
//				}
//			}
//			cur = next // 替换当前层
//		}
//		return 0
//	}
//	return bfs(start)
//}
//
//func inPath(node string, path []string) bool {
//	for _, word := range path {
//		if node == word {
//			return true
//		}
//	}
//	return false
//}
//
//func makeGraph(wordList []string) map[string][]string {
//	m := make(map[string][]string, len(wordList))
//	for _, word := range wordList {
//		m[word] = []string{}
//	}
//	for i := 0; i < len(wordList); i++ {
//		for j := i + 1; j < len(wordList); j++ {
//			if canTrans(wordList[i], wordList[j]) {
//				m[wordList[i]] = append(m[wordList[i]], wordList[j])
//				m[wordList[j]] = append(m[wordList[j]], wordList[i])
//			}
//		}
//	}
//	return m
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
//
//func main() {
//	beginWord := "qa"
//	endWord := "sq"
//	//wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
//	//wordList := []string{"hot", "dot", "dog", "lot", "log"}
//	wordList := []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}
//	min := ladderLength(beginWord, endWord, wordList)
//	fmt.Printf("%d\n", min)
//}
