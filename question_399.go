package leetcode

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	result := []float64{}

	// 构造图
	edges := map[string]map[string]float64{}
	for i, eq := range equations {
		A, B := eq[0], eq[1]
		val, rval := values[i], float64(1)/values[i]

		// 添加正向边
		if e, ok := edges[A]; ok {
			e[B] = val
		} else {
			edges[A] = map[string]float64{B: val}
		}

		// 添加反向边
		if e, ok := edges[B]; ok {
			e[A] = rval
		} else {
			edges[B] = map[string]float64{A: rval}
		}
	}

	for _, q := range queries {
		from, to := q[0], q[1]
		// 有不存在的元素
		if _, ok := edges[from]; !ok {
			result = append(result, float64(-1.0))
			continue
		}
		if _, ok := edges[to]; !ok {
			result = append(result, float64(-1.0))
			continue
		}
		// 除数被除数相等
		if from == to {
			result = append(result, float64(1.0))
			continue
		}
		// 根据路径计算结果
		_, ans := bfs(edges, from, to)
		result = append(result, ans)
	}

	return result
}

func bfs(edges map[string]map[string]float64, from, to string) (bool, float64) {
	// 记录已经访问过的节点
	visited := map[string]bool{}

	var level func(clevel []string, cvalues []float64) ([]string, []float64)
	level = func(clevel []string, cvalues []float64) ([]string, []float64) {
		// 没有可以访问的节点了
		if len(clevel) == 0 {
			return nil, nil
		}

		nlevel := []string{}
		nvalues := []float64{}

		for i, node := range clevel {
			val := cvalues[i]

			tos := edges[node]
			for to, v := range tos {
				// 已经访问过的节点不再处理
				if _, ok := visited[to]; !ok {
					nlevel = append(nlevel, to)
					nvalues = append(nvalues, val*v)
					visited[to] = true
				}
			}
		}
		if len(nlevel) == 0 {
			return nil, nil
		}
		return nlevel, nvalues
	}

	// 构造起点层
	ns := []string{from}
	vs := []float64{float64(1.0)}
	visited[from] = true
	// 逐层遍历
	for ns != nil {
		// 判断当前层是否到达目标
		for i, n := range ns {
			if n == to {
				return true, vs[i]
			}
		}
		ns, vs = level(ns, vs)
	}

	// 没有找到路径
	return false, float64(-1.0)
}

//func main() {
//	equations := [][]string{
//		{"a", "b"},
//	}
//	values := []float64{float64(0.5)}
//	queries := [][]string{
//		{"a", "b"},
//		{"b", "a"},
//		{"a", "c"},
//		{"x", "y"},
//	}
//	result := calcEquation(equations, values, queries)
//	fmt.Printf("%+v", result)
//}
