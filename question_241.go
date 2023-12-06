package leetcode

import (
	"strconv"
)

// 网友高赞题解 感觉比官方题解容易理解
func diffWaysToCompute(expression string) []int {
	// 如果是数字，直接返回
	if isDigit241(expression) {
		tmp, _ := strconv.Atoi(expression)
		return []int{tmp}
	}

	// 空切片
	var res []int
	// 遍历字符串
	for index, c := range expression {
		tmpC := string(c)
		if tmpC == "+" || tmpC == "-" || tmpC == "*" {
			// 如果是运算符，则计算左右两边的算式
			left := diffWaysToCompute(expression[:index])
			right := diffWaysToCompute(expression[index+1:])

			for _, leftNum := range left {
				for _, rightNum := range right {
					var addNum int
					if tmpC == "+" {
						addNum = leftNum + rightNum
					} else if tmpC == "-" {
						addNum = leftNum - rightNum
					} else {
						addNum = leftNum * rightNum
					}
					res = append(res, addNum)
				}
			}
		}
	}

	return res
}

func isDigit241(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return true
}
