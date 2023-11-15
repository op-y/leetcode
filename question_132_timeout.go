package leetcode

//import (
//	"math"
//)
//
//func minCut(s string) int {
//	l := len(s)
//	if l == 1 {
//		return 0
//	}
//
//	min := math.MaxInt
//
//	pos := 0
//	seg := 0
//
//	var f func(int)
//	f = func(start int) {
//		end := start + 1
//		if end == l+1 {
//			if pos == l {
//				if seg < min {
//					min = sig
//				}
//			}
//			return
//		}
//		for end <= l {
//			if ok := judge132(s[start:end]); ok {
//				seg++
//				pos = end
//				f(pos)
//				pos = start
//				seg--
//			}
//			end++
//		}
//	}
//
//	f(0)
//	return min
//}
//
//func judge132(s string) bool {
//	l := len(s)
//	if l == 1 {
//		return true
//	}
//	i, j := 0, l-1
//	for i < j {
//		if s[i] != s[j] {
//			return false
//		}
//		i++
//		j--
//	}
//	return true
//}
