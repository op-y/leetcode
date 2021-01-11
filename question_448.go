package leetcode

import (
    "fmt"
)

func findDisappearedNumbers(nums []int) []int {
	if nums == nil {
		return nil
	}
	length := len(nums)
	if length == 0 {
		return nil
	}

	idx := 0
	for idx < length {
		v := nums[idx]
        if v == 0 {
            idx++
            continue
        }
		if v == length {
			if v != nums[0] {
				tmp := nums[0]
				nums[0] = v
				nums[idx] = tmp
			} else {
                if idx == 0 {
				    idx++
                } else {
				    nums[idx] = 0
				    idx++
                }
			}
		} else {
			if v != nums[v] {
				tmp := nums[v]
				nums[v] = v
				nums[idx] = tmp
			} else {
                if idx == v {
				    idx++
                } else {
				    nums[idx] = 0
				    idx++
                }
			}
		}
        fmt.Printf("%v\n", nums)
	}

	result := make([]int, 0)
	for i, v := range nums {
		if i == 0 && v != length {
			result = append(result, length)
			continue
		}
		if v == 0 {
			result = append(result, i)
		}
	}
	return result
}
