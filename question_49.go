package leetcode

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	marks := []string{}

	l := len(strs)
	for i := 0; i < l; i++ {
		mark := sortString(strs[i])
		if len(marks) == 0 {
			marks = append(marks, mark)
			result = append(result, []string{strs[i]})
			continue
		}

		found := false
		for idx, item := range marks {
			if mark == item {
				result[idx] = append(result[idx], strs[i])
				found = true
				break
			}
		}
		if !found {
			marks = append(marks, mark)
			result = append(result, []string{strs[i]})
		}
	}

	return result
}

func sortString(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

//func main() {
//	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
//	r := groupAnagrams(strs)
//	fmt.Printf("%v\n", r)
//}
