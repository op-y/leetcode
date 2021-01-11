package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    if nums1 == nil {
       return nil
    }
    length := len(nums1)
    if length == 0 {
       return nil
    }

    result := make([]int, length)
    for idx, num1 := range nums1 {
        ok := false
        find := false
        for _, num2 := range nums2 {
            if num2 == num1 && !find {
                find = true
                continue
            }
            if num2 > num1 && find {
                ok = true
                result[idx] = num2
                break
            }
        }
        if !ok {
            result[idx] = -1
        }
    }
    return result
}
