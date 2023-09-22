package leetcode

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	set := make(map[int]bool, 0)
	set[n] = true
	for n != 1 {
		m := n
		sum := 0
		for m != 0 {
			mod := m % 10
			sum += mod * mod
			m /= 10
		}
		if ok, _ := set[sum]; ok {
			return false
		} else {
			set[sum] = true
		}
		n = sum
	}
	return true
}
