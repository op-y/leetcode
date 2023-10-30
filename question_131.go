package leetcode

func partition(s string) [][]string {
	ans := [][]string{}
	l := len(s)
	if l == 1 {
		return [][]string{{s}}
	}

	subs := []string{}
	pos := 0
	var f func(int)
	f = func(start int) {
		end := start + 1
		if end == l+1 {
			if pos == l {
				one := make([]string, len(subs))
				copy(one, subs)
				ans = append(ans, one)
			}
			return
		}
		for end <= l {
			if ok := judge131(s[start:end]); ok {
				subs = append(subs, s[start:end])
				pos = end
				f(pos)
				subs = subs[:len(subs)-1]
				pos = start
			}
			end++
		}
	}
	f(0)
	return ans
}

func judge131(s string) bool {
	l := len(s)
	if l == 1 {
		return true
	}
	i, j := 0, l-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
