package leetcode

func isPalindrome125(s string) bool {
	ss := []byte(s)
	length := len(ss)
	if length == 0 || length == 1 {
		return true
	}
	i := 0
	j := length - 1
	for i <= j {
		if !isAlphaBeta(uint8(ss[i])) {
			i++
			continue
		}
		if !isAlphaBeta(uint8(ss[j])) {
			j--
			continue
		}
		if !isCharEqual(uint8(ss[i]), uint8(ss[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphaBeta(c uint8) bool {
	if c >= uint8(48) && c <= uint8(57) {
		return true
	}
	if c >= uint8(65) && c <= uint8(90) {
		return true
	}
	if c >= uint8(97) && c <= uint8(122) {
		return true
	}
	return false
}

func isCharEqual(c1, c2 uint8) bool {
	if c1 >= uint8(65) && c1 <= uint8(90) {
		c1 += 32
	}
	if c2 >= uint8(65) && c2 <= uint8(90) {
		c2 += 32
	}
	if c1 == c2 {
		return true
	} else {
		return false
	}
}
