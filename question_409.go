package leetcode

func longestPalindrome409(s string) int {
	total := 0
	odd := false
	dict := make(map[byte]int, 0)
	ss := []byte(s)
	for _, c := range ss {
		if v, ok := dict[c]; ok {
			dict[c] = v + 1
		} else {
			dict[c] = 1
		}
	}

	for _, v := range dict {
		if v%2 == 0 {
			total += v
		} else {
			odd = true
			total += v - 1
		}
	}

	if odd {
		total += 1
	}

	return total
}
