import (
    "strings"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}
	c := s[0]
	sign := 1
	if c == byte('-') {
		sign = -1
		s = s[1:]
	} else if c == byte('+') {
		s = s[1:]
	}
	s = strings.TrimLeft(s, "0")
	digits := []int{}
	for i := 0; i < len(s); i++ {
		if !isDigit(s[i]) {
			break
		}
		digits = append(digits, int(s[i])-int(byte('0')))
	}
	sum := 0
	factor := 1
	for i := len(digits) - 1; i >= 0; i-- {
		nsum := sum + digits[i]*factor
		if sign == -1 && nsum > (1<<31) {
			sum = 1<<31
			break
		}
		if sign == 1 && nsum > (1<<31)-1 {
			sum = (1<<31)-1
			break
		}
		sum = nsum
		factor *= 10
        if factor < 0 {
		    if sign == -1 {
			    sum = 1<<31
			    break
		    }
		    if sign == 1 {
			    sum = (1<<31)-1
			    break
		    }
        }
	}
	sum *= sign
	return sum
}

func isDigit(c byte) bool {
	if c >= byte('0') && c <= byte('9') {
		return true
	} else {
		return false
	}
}
