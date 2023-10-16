package leetcode

func removeDuplicates80(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 1
	}

	pos := 0
	ch := nums[0]
	count := 0
	for i := 0; i < l; i++ {
		if nums[i] == ch {
			count++
			continue
		}
		if count == 1 {
			nums[pos] = ch
			pos++
		} else if count >= 2 {
			nums[pos], nums[pos+1] = ch, ch
			pos += 2
		}
		ch = nums[i]
		count = 1
	}
	if count == 1 {
		nums[pos] = ch
		pos++
	} else if count >= 2 {
		nums[pos], nums[pos+1] = ch, ch
		pos += 2
	}

	return pos
}
