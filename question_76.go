package leetcode

func minWindow(s string, t string) string {
	ls := len(s)
	lt := len(t)
	if ls < lt {
		return ""
	}

	min := ""
	target := genDict(t)
	left, right := 0, lt-1
	source := genDict(s[left : right+1])

	for right < ls && right-left+1 >= lt {
		if eq76(source, target) {
			if min == "" {
				min = s[left : right+1]
			} else if right-left+1 < len(min) {
				min = s[left : right+1]
			}
			source[s[left]]--
			left++
		} else {
			right++
			if right < ls {
				source[s[right]]++
			}
		}
	}

	return min
}

func genDict(s string) map[byte]int {
	d := map[byte]int{}
	for i := 0; i < len(s); i++ {
		d[s[i]]++
	}
	return d
}

func eq76(source, target map[byte]int) bool {
	for tc, tv := range target {
		if sv, ok := source[tc]; !ok || sv < tv {
			return false
		}
	}
	return true
}
