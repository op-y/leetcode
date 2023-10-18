package leetcode

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	result := [][]int{}
	result = append(result, []int{nums[0]})
	lastSets := [][]int{{nums[0]}}
	for i := 1; i < l; i++ {
		if nums[i] == nums[i-1] {
			curSets := [][]int{}
			for _, s := range lastSets {
				tmp := make([]int, len(s))
				copy(tmp, s)
				tmp = append(tmp, nums[i])
				curSets = append(curSets, tmp)
			}
			result = append(result, curSets...)
			lastSets = curSets
		} else {
			curSets := [][]int{}
			curSets = append(curSets, []int{nums[i]})
			for _, s := range result {
				tmp := make([]int, len(s))
				copy(tmp, s)
				tmp = append(tmp, nums[i])
				curSets = append(curSets, tmp)
			}
			result = append(result, curSets...)
			lastSets = curSets
		}
	}
	result = append(result, []int{})
	return result
}
