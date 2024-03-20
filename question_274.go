package leetcode

import (
	"sort"
)


func hIndex(citations []int) int {
	h := 0
	sort.Ints(citations)
	for i:=len(citations)-1; i>=0 && citations[i] > h; i-- {
		h++
	}
	return h
}

//func hIndex(citations []int) int {
//	l := len(citations)
//	counter := make([]int, l+1)
//	for _, v := range citations {
//		if v >= l {
//			counter[l]++
//		} else {
//			counter[v]++
//		}
//	}
//	for i, t := n, 0; i>=0; i-- {
//		t += counter[i]
//		if t >= i {
//			return i
//		}
//	}
//	return 0
//}
