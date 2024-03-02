package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int, numCourses) // 边 第一个维度是前置课程 第二个维度是前置课程->后续课程的边
		visited = make([]int, numCourses) // 标记节点状态 0未搜索 1在搜索过程中 2已经搜索过
		result []int // 保存拓扑排序结果 深度优先遍历 后序加入
		valid = true // 是否有环标记
		dfs func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			// v 没有搜索过递归dfs
			if visited[v] == 0 {
				dfs(v)
				// 检测到有向环直接返回
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				// 在搜索过程中遇到正在被处理的节点表示存在有向环
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	// 通过课程依赖关系构造边数据
	for _, vw := range prerequisites {
		edges[vw[1]] = append(edges[vw[1]], vw[0])
	}

	// 对所有节点做深度优先遍历
	for i:=0; i<numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	return valid
}
