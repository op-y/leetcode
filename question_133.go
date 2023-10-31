package leetcode

func cloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}

	processed := map[*GraphNode]bool{}
	cloneMap := map[*GraphNode]*GraphNode{}

	var bfs func(*GraphNode)
	bfs = func(node *GraphNode) {
		if _, ok := processed[node]; !ok {
			var cloneGraphNode *GraphNode
			if n, ok := cloneMap[node]; ok {
				// 获取未处理过的新节点
				cloneGraphNode = n
			} else {
				// 创建对应新节点
				cloneGraphNode = &GraphNode{Val: node.Val, Neighbors: []*GraphNode{}}
				cloneMap[node] = cloneGraphNode
			}

			for _, neighbor := range node.Neighbors {
				if cn, ok := cloneMap[neighbor]; ok {
					// 已经创建过的邻节点直接添加
					cloneGraphNode.Neighbors = append(cloneGraphNode.Neighbors, cn)
				} else {
					// 不存在的邻节点先创建再添加
					cloneNeighbor := &GraphNode{Val: neighbor.Val, Neighbors: []*GraphNode{}}
					cloneMap[neighbor] = cloneNeighbor
					cloneGraphNode.Neighbors = append(cloneGraphNode.Neighbors, cloneNeighbor)
				}
			}
			processed[node] = true
			for _, neighbor := range node.Neighbors {
				bfs(neighbor)
			}
		}
		return
	}
	bfs(node)
	return cloneMap[node]
}

//func main() {
//	n4 := &GraphNode{Val: 4, Neighbors: []*GraphNode{}}
//	n3 := &GraphNode{Val: 3, Neighbors: []*GraphNode{}}
//	n2 := &GraphNode{Val: 2, Neighbors: []*GraphNode{}}
//	n1 := &GraphNode{Val: 1, Neighbors: []*GraphNode{}}
//	n1.Neighbors = append(n1.Neighbors, n2, n4)
//	n2.Neighbors = append(n2.Neighbors, n1, n3)
//	n3.Neighbors = append(n3.Neighbors, n2, n4)
//	n4.Neighbors = append(n4.Neighbors, n1, n3)
//	cloneGraph(n1)
//}
