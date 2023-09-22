package leetcode

func findTheDifference(s string, t string) byte {
	ss := []byte(s)
	ts := []byte(t)

	if len(ss) == 0 {
		return ts[0]
	}

	dict := make(map[byte]int, 0)

	for _, b := range ss {
		if c, ok := dict[b]; ok {
			dict[b] = c + 1
		} else {
			dict[b] = 1
		}
	}

	for _, b := range ts {
		if c, ok := dict[b]; ok {
			dict[b] = c - 1
		} else {
			return b
		}
	}

	target := byte(' ')
	for b, c := range dict {
		if c == -1 {
			target = b
		}
	}
	return target
}
