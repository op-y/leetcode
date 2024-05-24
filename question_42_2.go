package leetcode

func trap2(height []int) int {
	ans := 0
	stack := []int{}
	for i, h := range height {
		// 当前元素高度大于栈顶元素 可能存在一个收集区间
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// left 不存在
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			cWidth := i - left - 1
			cHeight := min(height[left], h) - height[top]
			ans += cWidth * cHeight
		}
		stack = append(stack, i)
	}

	return ans
}
