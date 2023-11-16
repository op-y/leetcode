package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 结束词不在图中
	if !inPath(endWord, wordList) {
		return 0
	}
	// 将输入节点一起放入构造图
	if !inPath(beginWord, wordList) {
		wordList = append(wordList, beginWord)
	}
	graph := makeGraph(wordList)

	paths := shortestPath(beginWord, endWord, graph)
	if len(paths) > 0 { // 存在路径
		return len(paths[0])
	}
	return 0
}

func shortestPath(start, end string, graph map[string][]string) [][]string {
	paths := [][]string{}

	path := []string{start}
	var dfs func(string)
	dfs = func(node string) {
		if node == end {
			// 找到一条路径
			one := make([]string, len(path))
			copy(one, path)
			paths = append(paths, one)
		}
		for _, next := range graph[node] {
			// 路径不能成环
			if !inPath(next, path) {
				path = append(path, next)
				dfs(next)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(start)

	// 没有路径
	if len(paths) == 0 {
		return paths
	}

	shortest := len(paths[0])
	for _, path := range paths {
		if len(path) < shortest {
			shortest = len(path)
		}
	}
	ans := [][]string{}
	for _, path := range paths {
		if len(path) == shortest {
			ans = append(ans, path)
		}
	}
	return ans
}

func inPath(node string, path []string) bool {
	for _, word := range path {
		if node == word {
			return true
		}
	}
	return false
}

func makeGraph(wordList []string) map[string][]string {
	m := make(map[string][]string, len(wordList))
	for _, word := range wordList {
		m[word] = []string{}
	}
	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if canTrans(wordList[i], wordList[j]) {
				m[wordList[i]] = append(m[wordList[i]], wordList[j])
				m[wordList[j]] = append(m[wordList[j]], wordList[i])
			}
		}
	}
	return m
}

func canTrans(from, to string) bool {
	l := len(from)
	found := false
	for i := 0; i < l; i++ {
		if from[i] == to[i] {
			continue
		}
		if !found {
			found = true
		} else {
			// 超过一个字符的差异
			return false
		}
	}
	return found
}
