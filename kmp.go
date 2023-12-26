package leetcode

func initNexts(pat string) []int {
	m := len(pat)
	nexts := make([]int, m)

	for j, i := 0, 1; i < m; i++ {
		for j > 0 && pat[i] != pat[j] {
			j = nexts[j-1]
		}
		if pat[i] == pat[j] {
			j++
			nexts[i] = j
		}
	}

	return nexts
}

func kmp(txt, pat string) int {
	nexts := initNexts(pat)
	n, m := len(txt), len(pat)
	i, j := 0, 0
	for i < n && j < m {
		if txt[i] == pat[j] {
			i++
			j++
			continue
		}
		if j == 0 {
			i++
		} else {
			j = nexts[j-1]
		}
	}
	if j == m {
		return i - m
	}
	return -1
}

//func main() {
//	//txt := "AABRAACADABRAACAADABRA"
//	//pat := "AACAA"
//	txt := "aabaabaaf"
//	pat := "aabaaf"
//	r := kmp(txt, pat)
//	fmt.Println(r)
//}
