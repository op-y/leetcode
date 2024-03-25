package leetcode

func minMutation(startGene string, endGene string, bank []string) int {
	edges := map[string][]string{}

	// start一定合法 放入bank中
	bank = append(bank, startGene)

	// 构造图
	n := len(bank)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			can := canTrans(bank[i], bank[j])
			if can {
				if _, ok := edges[bank[i]]; ok {
					edges[bank[i]] = append(edges[bank[i]], bank[j])
				} else {
					edges[bank[i]] = []string{bank[j]}
				}

				if _, ok := edges[bank[j]]; ok {
					edges[bank[j]] = append(edges[bank[j]], bank[i])
				} else {
					edges[bank[j]] = []string{bank[i]}
				}
			}
		}
	}

	// 判断end是否在bank中
	if _, ok := edges[endGene]; !ok {
		return -1
	}

	// bfs 遍历图尝试到达end
	count := 0
	q := []string{startGene}
	visited := map[string]bool{startGene: true}
	for len(q) > 0 {
		next := []string{}
		count++
		for _, gene := range q {
			for _, g := range edges[gene] {
				if g == endGene {
					return count
				}
				if _, ok := visited[g]; !ok {
					visited[g] = true
					next = append(next, g)
				}
			}
		}
		q = next
	}
	return -1
}

func canTrans(from, to string) bool {
	diff := 0
	for i := 0; i < 8; i++ {
		if from[i] != to[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	if diff == 1 {
		return true
	}
	return false
}

//func main() {
//	startGene := "AACCGGTT"
//	endGene := "AAACGGTA"
//	bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
//	fmt.Println(minMutation(startGene, endGene, bank))
//}
