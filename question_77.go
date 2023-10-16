package leetcode

func combine(n int, k int) (ans [][]int) {
	elems := []int{}
	selected := make([]bool, n)
	for i := 0; i < n; i++ {
		elems = append(elems, i+1)
	}

	current := []int{}

	var pick func(int)
	pick = func(pos int) {
		if len(current) == k {
			one := make([]int, k)
			copy(one, current)
			ans = append(ans, one)
			return
		}
		for i := pos; i < n; i++ {
			if selected[i] {
				continue
			}
			selected[i] = true
			current = append(current, elems[i])
			pick(i + 1)
			selected[i] = false
			current = current[0 : len(current)-1]
		}
	}
	pick(0)
	return
}
