package leetcode

func candy(ratings []int) int {
	ans := 0
	l := len(ratings)
	// from left to right
	left := make([]int, l)
	for i, rating := range ratings {
		if i > 0 && rating > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	// from right to left
	right := 0
	for i := l - 1; i >= 0; i-- {
		if i < l-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return ans
}
