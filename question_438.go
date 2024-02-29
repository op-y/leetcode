package leetcode

func findAnagrams(s string, p string) []int {
	result := []int{}
	ls, lp := len(s), len(p)

	if lp > ls {
		return result
	}

	// p dict
	dictp := map[byte]int{}
	for i := 0; i < lp; i++ {
		c := p[i]
		dictp[c]++
	}

	// s dict
	dicts := map[byte]int{}
	for i := 0; i < ls; i++ {
		// 先计算lp-1个字符
		c := s[i]
		if i < lp-1 {
			dicts[c]++
			continue
		}

		dicts[c]++
		if dictEqual(dicts, dictp) {
			result = append(result, i-lp+1)
		}

		start := s[i-lp+1]
		dicts[start]--
		if dicts[start] == 0 {
			delete(dicts, start)
		}
	}

	return result
}

func dictEqual(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}
	for ka, va := range a {
		vb, ok := b[ka]
		if !ok {
			return false
		}
		if va != vb {
			return false
		}
	}
	return true
}

//func main() {
//	s := "abab"
//	p := "ab"
//	r := findAnagrams(s, p)
//	fmt.Println(r)
//}
