package leetcode

func trailingZeroes(n int) int {
	result := 0
	factor := 5
	i := n / factor
	for ; i != 0; factor *= 5 {
		result += i
	}
	return result
}
