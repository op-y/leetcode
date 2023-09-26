package leetcode

func longestValidParentheses(s string) int {
	l := len(s)
	if l < 2 {
		return 0
	}

	max := 0
	length := 0
	stack := []int{-1}
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			idx := len(stack) - 1
			stack = stack[0:idx]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				idx = len(stack) - 1
				length = i - stack[idx]
				if length > max {
					max = length
				}
			}
		}
	}

	return max
}

//func main() {
//	s := "(()()(())(("
//	max := longestValidParentheses(s)
//	fmt.Printf("%d\n", max)
//}
