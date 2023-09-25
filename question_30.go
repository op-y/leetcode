package leetcode

func findSubstring(s string, words []string) []int {
	result := []int{}

	ls := len(s)
	lw := len(words[0])
	wcnt := len(words)

	if ls < wcnt*lw {
		return []int{}
	}

	wordmap := make(map[string]int)
	for _, word := range words {
		if cnt, ok := wordmap[word]; ok {
			wordmap[word] = cnt + 1
		} else {
			wordmap[word] = 1
		}
	}

	tlen := wcnt * lw
	for i := 0; ls-i >= tlen; i++ {
		if cmpWordsMap(wordmap, s, i, lw, wcnt, tlen) {
			result = append(result, i)
		}
	}

	return result
}

func cmpWordsMap(wordmap map[string]int, s string, i, lw, wcnt, tlen int) bool {
	hit := 0
	wordmap2 := make(map[string]int)
	for word, cnt := range wordmap {
		wordmap2[word] = cnt
	}
	start := i
	end := i + lw
	for end-i <= tlen {
		word := s[start:end]
		if cnt, ok := wordmap2[word]; ok {
			if cnt <= 0 {
				return false
			} else {
				wordmap2[word] -= 1
				hit++
				if hit == wcnt {
					return true
				}
			}
		} else {
			return false
		}
		start += lw
		end += lw
	}
	return false
}

//func main() {
//	s := "barfoothefoobarman"
//	words := []string{"foo", "bar"}
//
//	result := findSubstring(s, words)
//	fmt.Printf("%+v\n", result)
//}
