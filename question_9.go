package leetcode

func isPalindrome9(x int) bool {
	origin := x
	if x < 0 {
		return false
	}
	// 可以只反转一半后做比较
	result := 0
	for {
		remainder := x % 10
		x = x / 10
		if x > 0 {
			result = (result + remainder) * 10
		} else {
			result = result + remainder
			break
		}
	}
	if result == origin {
		return true
	} else {
		return false
	}
}
