package leetcode

func firstMissingPositive(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 1
	}
	if l == 1 && nums[0] != 1 {
		return 1
	}

	head := 0
	tail := l - 1
	for head <= tail {
		if nums[head] <= 0 && nums[tail] > 0 {
			nums[head], nums[tail] = nums[tail], nums[head]
			tail--
			continue
		}
		num := nums[head]
		if num > 0 && num-1 < l && num-1 <= tail && nums[num-1] != num {
			nums[head], nums[num-1] = nums[num-1], nums[head]
			continue
		}
		head++
	}

	for i := 0; i < l; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return l + 1
}

//func main() {
//	nums := []int{2, 1, 1, 0, 3, 0, 4, 3, 2}
//	result := firstMissingPositive(nums)
//	fmt.Printf("%d\n", result)
//}
