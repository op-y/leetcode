package leetcode

func intersection(nums1 []int, nums2 []int) []int {
    if nums1 == nil || nums2 == nil {
        return nil
    }

    items := make(map[int]bool, 0)
    result := make([]int, 0)

    for _, v :=  range nums1 {
        if _, ok := items[v]; ok {
            continue
        } else {
            items[v] = false
        }
    }

    for _, v := range nums2 {
        if processed, ok := items[v]; ok {
            if !processed {
                result = append(result, v)
                items[v] = true
            }
        }
    }

    return result
}
