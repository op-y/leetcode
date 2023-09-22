package leetcode

func isPalindrome9(x int) bool {
	origin := x
	if x < 0 {
		return false
	}
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
