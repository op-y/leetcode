package leetcode

func subsets(nums []int) [][]int {
	result := [][]int{{}}
	set := find78(nums, 0)
	result = append(result, set...)
	return result
}

func find78(nums []int, i int) [][]int {
	tmp := [][]int{}
	tmp = append(tmp, []int{nums[i]})
	if i+1 < len(nums) {
		ntmp := find78(nums, i+1)
		for _, item := range ntmp {
			nitem := []int{nums[i]}
			nitem = append(nitem, item...)
			tmp = append(tmp, nitem)
		}
		for _, item := range ntmp {
			tmp = append(tmp, item)
		}
	}
	return tmp
}
