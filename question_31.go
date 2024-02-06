package leetcode

func nextPermutation(nums []int) {
	l := len(nums)
	if l == 1 {
		return
	}

	for i := l - 2; i >= 0; i-- {
		ok := checkTail(nums, i)
		if !ok {
			continue
		} else {
			sort31(nums, i+1)
			return
		}
	}
	sort31(nums, 0)
}

// 这个处理不是最简单的 可以从后往前找第一个非降序元素
func checkTail(nums []int, idx int) bool {
	pos := -1
	value := 0
	l := len(nums)
	for i := l - 1; i > idx; i-- {
		if nums[i] > nums[idx] && pos < 0 {
			value = nums[i]
			pos = i
			continue
		}
		if nums[i] > nums[idx] && nums[i] < value && pos > 0 {
			pos = i
			value = nums[i]
			continue
		}
	}

	if pos > 0 {
		nums[idx], nums[pos] = nums[pos], nums[idx]
		return true
	}
	return false
}

func sort31(nums []int, pos int) {
	l := len(nums)
	for i := pos; i < l-1; i++ {
		for j := l - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

//func main() {
//	nums := []int{1, 1, 5}
//	nextPermutation(nums)
//	fmt.Printf("%+v\n", nums)
//}
