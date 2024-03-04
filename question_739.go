package leetcode

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))

	stack := []int{} // 单调递减栈
	for i:=0; i<len(temperatures); i++ {
		// 栈为空
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		// 当前遍历温度大于栈顶位置温度
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = i-stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}
	
	// 处理剩余元素
	for i:=0; i<len(stack); i++ {
		ans[stack[i]] = 0
	}	

	return ans
}
