package leetcode

func canJump(nums []int) bool {
	l := len(nums)
	if l == 1 {
		return true
	}

	steps := make([]int, l)
	for i := 1; i < l; i++ {
		steps[i] = -1
	}

	for i := 0; i < l-1; i++ {
		// 当前位置无法到达
		if steps[i] == -1 {
			return false
		}

		v := nums[i]
		if v <= 0 {
			continue
		}
		for j := 1; j <= v; j++ {
			idx := i + j
			if idx < l && (steps[idx] == -1 || (steps[i]+1) < steps[idx]) {
				steps[idx] = steps[i] + 1
			}
		}
	}
	if steps[l-1] == -1 {
		return false
	} else {
		return true
	}
}

//func main() {
//	nums := []int{3, 2, 1, 0, 4, 5}
//	ok := canJump(nums)
//	fmt.Printf("%v\n", ok)
//}
