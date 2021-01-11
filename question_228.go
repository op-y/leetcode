package leetcode

import (
    "fmt"
)

func summaryRanges(nums []int) []string {
    if nums == nil {
        return nil
    }
    length := len(nums)
    if length == 0 {
        return nil
    }
    result := make([]string, 0)
    start := 0
    end := 0
    for i:=0; i<length; i++ {
        if 1 == nums[i] - nums[end]  {
            end = i
        } else if nums[i] - nums[end] > 1 {
            if end > start {
                item := fmt.Sprintf("%d->%d", nums[start], nums[end])
                result = append(result, item)
                start = i
                end = i
            } else {
                item := fmt.Sprintf("%d", nums[start])
                result = append(result, item)
                start = i
                end = i
            }
        }
    }
    if end > start {
        item := fmt.Sprintf("%d->%d", nums[start], nums[end])
        result = append(result, item)
    } else {
        item := fmt.Sprintf("%d", nums[start])
        result = append(result, item)
    }
    return result
}
