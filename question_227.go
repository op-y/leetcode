package leetcode

import (
	"strconv"
)

func calculate(s string) int {
	l := len(s)
	nums := []int{}
	ops := []byte{}

	var getNum func(int) (int, int)
	getNum = func(x int) (int, int) {
		bs := []byte{}
		for x < l && s[x] == ' ' {
			x++
		}
		for x < l && s[x]-'0' >= 0 && s[x]-'0' <= 9 {
			bs = append(bs, s[x])
			x++
		}
		n, _ := strconv.Atoi(string(bs))
		return n, x
	}

	i := 0
	for i < l {
		// 忽略空格
		if s[i] == ' ' {
			i++
			continue
		}

		// 加减法优先级低先入栈
		if s[i] == '+' || s[i] == '-' {
			ops = append(ops, s[i])
			i++
			continue
		}

		// 乘除法情况
		if s[i] == '*' {
			i++

			n1 := nums[len(nums)-1]
			nums = nums[:len(nums)-1]

			n2, next := getNum(i)
			i = next

			n := n1 * n2
			nums = append(nums, n)
			continue
		}
		if s[i] == '/' {
			i++

			n1 := nums[len(nums)-1]
			nums = nums[:len(nums)-1]

			n2, next := getNum(i)
			i = next

			n := n1 / n2
			nums = append(nums, n)
			continue
		}

		// 一个数字的情况
		n, next := getNum(i)
		nums = append(nums, n)
		i = next
	}

	for len(ops) > 0 {
		op := ops[0]
		ops = ops[1:]
		n1 := nums[0]
		n2 := nums[1]
		if op == '+' {
			n := n1 + n2
			nums[1] = n
			nums = nums[1:]
			continue
		}
		if op == '-' {
			n := n1 - n2
			nums[1] = n
			nums = nums[1:]
			continue
		}
	}

	return nums[0]
}

//func main() {
//	s := "1-1+1"
//	fmt.Println(calculate(s))
//}
